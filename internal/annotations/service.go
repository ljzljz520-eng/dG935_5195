package annotations

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }

// Save persists a new version of an annotation for doc. Versioning happens
// atomically inside a single bbolt transaction so each save lands under a
// distinct key and no prior version is overwritten (see AppendAnnotation).
func (s *Service) Save(doc, author, content string) (domain.Annotation, error) {
	return s.Store.AppendAnnotation(doc, author, content)
}
func (s *Service) Page(doc string, page, size int) ([]domain.Annotation, bool, error) {
	all, e := s.Store.ListAnnotations(doc)
	if e != nil {
		return nil, false, e
	}
	if page < 1 || size < 1 {
		return nil, false, domain.ErrInvalid
	}
	start := (page - 1) * size
	if start >= len(all) {
		return []domain.Annotation{}, false, nil
	}
	end := start + size
	more := true
	if end >= len(all) {
		end = len(all)
		more = false
	}
	return all[start:end], more, nil
}
func (s *Service) Latest(doc string) (domain.Annotation, error) {
	all, e := s.Store.ListAnnotations(doc)
	if e != nil || len(all) == 0 {
		return domain.Annotation{}, domain.ErrNotFound
	}
	return all[len(all)-1], nil
}
