package store

import (
	"accountingcollab/internal/domain"
	"path/filepath"
	"testing"
)

func TestPersistenceSurvivesReopen(t *testing.T) {
	p := filepath.Join(t.TempDir(), "data.db")
	s, e := Open(p)
	if e != nil {
		t.Fatal(e)
	}
	w := domain.NewWorkspace("w", "盘", "o")
	if e = s.PutWorkspace(w); e != nil {
		t.Fatal(e)
	}
	s.Close()
	s, e = Open(p)
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if _, e = s.GetWorkspace("w"); e != nil {
		t.Fatal(e)
	}
}
func TestStoreMissing(t *testing.T) {
	s, _ := Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	if _, e := s.GetDocument("x"); e == nil {
		t.Fatal("expected missing")
	}
}
