package export

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestExportSnapshot(t *testing.T) {
	s, _ := store.Open(filepath.Join(t.TempDir(), "x"))
	defer s.Close()
	s.PutDocument(domain.NewDocument("d", "w", "Doc"))
	s.PutAnnotation(domain.NewAnnotation("a", "d", "u", "text", 1))
	v, e := Build(s, "d")
	if e != nil || !IsAuditable(v) {
		t.Fatal(e)
	}
	if len(Render(v)) == 0 {
		t.Fatal()
	}
}
func TestEncode(t *testing.T) {
	b, e := Encode(Snapshot{})
	if e != nil || len(b) == 0 {
		t.Fatal(e)
	}
}
