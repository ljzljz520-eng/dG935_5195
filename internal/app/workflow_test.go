package app

import (
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestWorkflowAnnotationRoundTrip(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := New(s)
	w, _ := a.StartWorkspace("w", "盘", "o")
	d, _ := a.AddDocument("d", w.ID, "凭证")
	a.AddAnnotation(d.ID, "u", "one")
	a.AddAnnotation(d.ID, "u", "two")
	v, e := a.Export(d.ID)
	if e != nil || len(v.Annotations) != 2 {
		t.Fatal(e, len(v.Annotations))
	}
}
func TestWorkflowReviewCompletion(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := New(s)
	r, e := a.CompleteReview("d", "reviewer", "approved")
	if e != nil || r.Decision != "approved" {
		t.Fatal(e)
	}
}
func TestWorkflowExportSnapshot(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	a := New(s)
	a.StartWorkspace("w", "盘", "o")
	d, _ := a.AddDocument("d", "w", "凭证")
	a.AddAnnotation(d.ID, "u", "x")
	if _, e := a.Export(d.ID); e != nil {
		t.Fatal(e)
	}
}
