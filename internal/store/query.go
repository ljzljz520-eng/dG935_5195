package store

import (
	"accountingcollab/internal/domain"
	"encoding/json"
	"go.etcd.io/bbolt"
	"sort"
)

func (s *Store) PutAnnotation(v domain.Annotation) error {
	if e := domain.ValidateAnnotation(v); e != nil {
		return e
	}
	return put(s, buckets[2], v.ID, v)
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
