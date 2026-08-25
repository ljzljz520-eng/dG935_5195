package app

import (
	"accountingcollab/internal/store"
	"path/filepath"
	"testing"
)

func TestCLI(t *testing.T) {
	s, e := store.Open(filepath.Join(t.TempDir(), "x"))
	if e != nil {
		t.Fatal(e)
	}
	a := New(s)
	if _, e = a.StartWorkspace("w", "n", "o"); e != nil {
		t.Fatal(e)
	}
	a.Close()
}
