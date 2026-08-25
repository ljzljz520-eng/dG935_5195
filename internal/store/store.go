package store

import (
	"accountingcollab/internal/domain"
	"encoding/json"
	"go.etcd.io/bbolt"
	"path/filepath"
	"sync"
)

var buckets = [][]byte{[]byte("workspaces"), []byte("documents"), []byte("annotations"), []byte("reviews")}

type Store struct {
	db *bbolt.DB
	mu sync.RWMutex
}

func Open(path string) (*Store, error) {
	db, err := bbolt.Open(filepath.Clean(path), 0600, nil)
	if err != nil {
		return nil, err
	}
	s := &Store{db: db}
	err = db.Update(func(tx *bbolt.Tx) error {
		for _, b := range buckets {
			if _, e := tx.CreateBucketIfNotExists(b); e != nil {
				return e
			}
		}
		return nil
	})
	if err != nil {
		db.Close()
		return nil, err
	}
	return s, nil
}
func (s *Store) Close() error { s.mu.Lock(); defer s.mu.Unlock(); return s.db.Close() }
func put[T any](s *Store, b []byte, key string, v T) error {
	data, e := json.Marshal(v)
	if e != nil {
		return e
	}
	return s.db.Update(func(tx *bbolt.Tx) error { return tx.Bucket(b).Put([]byte(key), data) })
}
func get[T any](s *Store, b []byte, key string, out *T) error {
	return s.db.View(func(tx *bbolt.Tx) error {
		v := tx.Bucket(b).Get([]byte(key))
		if v == nil {
			return domain.ErrNotFound
		}
		return json.Unmarshal(v, out)
	})
}
func (s *Store) PutWorkspace(v domain.Workspace) error {
	if e := domain.ValidateWorkspace(v); e != nil {
		return e
	}
	return put(s, buckets[0], v.ID, v)
}
func (s *Store) GetWorkspace(id string) (domain.Workspace, error) {
	var v domain.Workspace
	e := get(s, buckets[0], id, &v)
	return v, e
}
func (s *Store) PutDocument(v domain.Document) error {
	if e := domain.ValidateDocument(v); e != nil {
		return e
	}
	return put(s, buckets[1], v.ID, v)
}
func (s *Store) GetDocument(id string) (domain.Document, error) {
	var v domain.Document
	e := get(s, buckets[1], id, &v)
	return v, e
}
func (s *Store) PutReview(v domain.ReviewTask) error {
	if e := domain.ValidateReview(v); e != nil {
		return e
	}
	return put(s, buckets[3], v.ID, v)
}
func (s *Store) GetReview(id string) (domain.ReviewTask, error) {
	var v domain.ReviewTask
	e := get(s, buckets[3], id, &v)
	return v, e
}
func (s *Store) RawPut(bucket, key string, val any) error {
	switch bucket {
	case "workspaces":
		return put(s, buckets[0], key, val)
	case "documents":
		return put(s, buckets[1], key, val)
	case "annotations":
		return put(s, buckets[2], key, val)
	case "reviews":
		return put(s, buckets[3], key, val)
	default:
		return domain.ErrInvalid
	}
}
