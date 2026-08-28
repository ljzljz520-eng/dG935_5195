package docs

import (
	"accountingcollab/internal/annotations"
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestAnnotationVersionsDistinct(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "annotations.db"))
	if e != nil {
		t.Fatal(e)
	}
	defer s.Close()
	if e = s.PutDocument(domain.NewDocument("doc", "workspace", "Ledger")); e != nil {
		t.Fatal(e)
	}
	svc := annotations.New(s)
	first, e := svc.Save("doc", "author", "one")
	if e != nil {
		t.Fatal(e)
	}
	second, e := svc.Save("doc", "author", "two")
	if e != nil {
		t.Fatal(e)
	}
	if first.Version == second.Version {
		t.Fatalf("annotation versions must be distinct: %d", first.Version)
	}
}
func TestGuideValidity(t *testing.T) {
	if !Valid(NewGuide("a", "b")) {
		t.Fatal("invalid guide")
	}
}
