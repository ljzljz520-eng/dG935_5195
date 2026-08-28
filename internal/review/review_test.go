package review

import (
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestReviewCompletion(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	svc := New(s)
	r, e := svc.Open("d", "reviewer")
	if e != nil {
		t.Fatal(e)
	}
	r, e = svc.Complete(r.ID, "approved")
	if e != nil || !svc.IsComplete(r) {
		t.Fatal(e)
	}
}
func TestReviewLabels(t *testing.T) {
	if Label("approved") == "" || !NeedsFollowup("changes_requested") {
		t.Fatal()
	}
	if !AllowedReviewer("ab") {
		t.Fatal()
	}
}
