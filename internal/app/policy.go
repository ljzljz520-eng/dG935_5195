package app

import "accountingcollab/internal/domain"

func EnsureTransition(d domain.Document, to string) error {
	if !domain.CanTransition(d.Status, to) {
		return domain.ErrConflict
	}
	d.Status = to
	return nil
}
func WorkflowReady(d domain.Document) bool { return d.ID != "" && d.Status != "archived" }
func RoleCanReview(role string) bool       { return role == "reviewer" || role == "owner" }
