package annotations

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestPageAnnotations(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.PutDocument(domain.NewDocument("d", "w", "Doc"))
	svc := New(s)
	svc.Save("d", "a", "one")
	svc.Save("d", "a", "two")
	got, _, e := svc.Page("d", 1, 2)
	if e != nil || len(got) != 2 {
		t.Fatal(got, e)
	}
}
func TestContentHelpers(t *testing.T) {
	if !IsSubstantive("abc") {
		t.Fatal()
	}
	if PageCount(5, 2) != 3 {
		t.Fatal()
	}
}
