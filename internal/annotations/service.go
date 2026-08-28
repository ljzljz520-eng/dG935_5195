package annotations

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"fmt"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Save(doc, author, content string) (domain.Annotation, error) {
	current, _ := s.Store.ListAnnotations(doc)
	version := 1
	if len(current) > 0 {
		version = current[len(current)-1].Version + 1
	}
	a := domain.NewAnnotation(fmt.Sprintf("%s-%d", doc, version), doc, author, content, 1)
	if e := s.Store.PutAnnotation(a); e != nil {
		return a, e
	}
	return a, nil
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
