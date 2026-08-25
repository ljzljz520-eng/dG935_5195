package store

import (
	"accountingcollab/internal/domain"
	"encoding/json"
	"fmt"
	"go.etcd.io/bbolt"
	"sort"
)

func (s *Store) PutAnnotation(v domain.Annotation) error {
	if e := domain.ValidateAnnotation(v); e != nil {
		return e
	}
	return put(s, buckets[2], v.ID, v)
}

// AppendAnnotation stores the next version of an annotation for doc inside a
// single read+write transaction. The version is derived from the highest
// existing version for that document and the ID is built from it, so every
// save lands under a distinct key (no overwriting of prior versions) and two
// concurrent saves cannot collide on the same version number.
func (s *Store) AppendAnnotation(doc, author, content string) (domain.Annotation, error) {
	if doc == "" || author == "" || content == "" {
		return domain.Annotation{}, domain.ErrInvalid
	}
	var a domain.Annotation
	e := s.db.Update(func(tx *bbolt.Tx) error {
		b := tx.Bucket(buckets[2])
		maxVersion := 0
		if e := b.ForEach(func(k, v []byte) error {
			var existing domain.Annotation
			if e := json.Unmarshal(v, &existing); e != nil {
				return e
			}
			if existing.DocumentID == doc && existing.Version > maxVersion {
				maxVersion = existing.Version
			}
			return nil
		}); e != nil {
			return e
		}
		version := maxVersion + 1
		a = domain.NewAnnotation(fmt.Sprintf("%s-%d", doc, version), doc, author, content, version)
		data, e := json.Marshal(a)
		if e != nil {
			return e
		}
		return b.Put([]byte(a.ID), data)
	})
	if e != nil {
		return domain.Annotation{}, e
	}
	return a, nil
}
func (s *Store) GetAnnotation(id string) (domain.Annotation, error) {
	var v domain.Annotation
	e := get(s, buckets[2], id, &v)
	return v, e
}
func (s *Store) ListAnnotations(doc string) ([]domain.Annotation, error) {
	out := []domain.Annotation{}
	e := s.db.View(func(tx *bbolt.Tx) error {
		return tx.Bucket(buckets[2]).ForEach(func(k, v []byte) error {
			var a domain.Annotation
			if e := json.Unmarshal(v, &a); e != nil {
				return e
			}
			if a.DocumentID == doc {
				out = append(out, a)
			}
			return nil
		})
	})
	sort.Slice(out, func(i, j int) bool { return out[i].Version < out[j].Version })
	return out, e
}
