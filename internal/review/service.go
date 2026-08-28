package review

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"time"
)

type Service struct{ Store *store.Store }

func New(s *store.Store) *Service { return &Service{Store: s} }
func (s *Service) Open(doc, reviewer string) (domain.ReviewTask, error) {
	r := domain.NewReviewTask("review-"+doc, doc, reviewer)
	if e := s.Store.PutReview(r); e != nil {
		return r, e
	}
	return r, nil
}
func (s *Service) Complete(id, decision string) (domain.ReviewTask, error) {
	r, e := s.Store.GetReview(id)
	if e != nil {
		return r, e
	}
	if !domain.ReviewDecisionValid(decision) {
		return r, domain.ErrInvalid
	}
	r.Decision = decision
	r.CompletedAt = time.Now().UTC()
	e = s.Store.PutReview(r)
	return r, e
}
func (s *Service) IsComplete(r domain.ReviewTask) bool {
	return r.CompletedAt.Unix() > 0 && domain.ReviewDecisionValid(r.Decision)
}
