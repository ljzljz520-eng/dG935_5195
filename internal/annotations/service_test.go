package annotations

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"os"
	"path/filepath"
	"testing"
)

// openStore opens a bbolt store at a temp path and returns it together with a
// closer. The path is returned so the same file can be reopened to verify that
// annotation history survives a process restart.
func openStore(t *testing.T, name string) (*store.Store, string) {
	t.Helper()
	path := filepath.Join(t.TempDir(), name)
	s, e := store.Open(path)
	if e != nil {
		t.Fatalf("open store: %v", e)
	}
	t.Cleanup(func() { _ = os.Remove(path) })
	return s, path
}

// TestSaveRetainsEveryVersion reproduces the reported defect: saving several
// annotations in a row used to overwrite earlier versions because every saved
// Version field was 1, so each save after the first produced the same key and
// clobbered the previous one. All saves must be retained and look-back-able.
func TestSaveRetainsEveryVersion(t *testing.T) {
	s, _ := openStore(t, "ann.db")
	defer s.Close()
	svc := New(s)

	contents := []string{"请核对金额", "金额无误，但日期有误", "已更正日期", "复核通过"}
	for i, c := range contents {
		a, e := svc.Save("doc-1", "owner", c)
		if e != nil {
			t.Fatalf("save %d: %v", i+1, e)
		}
		if a.Version != i+1 {
			t.Fatalf("save %d: version = %d, want %d", i+1, a.Version, i+1)
		}
		if a.Content != c {
			t.Fatalf("save %d: content = %q, want %q", i+1, a.Content, c)
		}
		if a.ID != "doc-1-1" && i == 0 {
			// first save id
		}
	}

	// Look back: paging must return every saved version in order.
	page, more, e := svc.Page("doc-1", 1, 100)
	if e != nil {
		t.Fatalf("page: %v", e)
	}
	if more {
		t.Fatalf("did not expect more pages")
	}
	if len(page) != len(contents) {
		t.Fatalf("retained %d versions, want %d", len(page), len(contents))
	}
	for i, a := range page {
		if a.Version != i+1 {
			t.Fatalf("page[%d].Version = %d, want %d", i, a.Version, i+1)
		}
		if a.Content != contents[i] {
			t.Fatalf("page[%d].Content = %q, want %q", i, a.Content, contents[i])
		}
	}

	// Latest must point at the most recent save.
	latest, e := svc.Latest("doc-1")
	if e != nil {
		t.Fatalf("latest: %v", e)
	}
	if latest.Content != contents[len(contents)-1] {
		t.Fatalf("latest content = %q, want %q", latest.Content, contents[len(contents)-1])
	}
}

// TestSaveAcrossReopen verifies that saved annotation history survives a
// store close and reopen (process restart), matching the persistence contract.
func TestSaveAcrossReopen(t *testing.T) {
	s, path := openStore(t, "reopen.db")
	svc := New(s)
	if _, e := svc.Save("doc-1", "owner", "第一版"); e != nil {
		t.Fatalf("save 1: %v", e)
	}
	if _, e := svc.Save("doc-1", "owner", "第二版"); e != nil {
		t.Fatalf("save 2: %v", e)
	}
	if e := s.Close(); e != nil {
		t.Fatalf("close: %v", e)
	}

	s2, e := store.Open(path)
	if e != nil {
		t.Fatalf("reopen: %v", e)
	}
	defer s2.Close()
	svc2 := New(s2)

	page, _, e := svc2.Page("doc-1", 1, 100)
	if e != nil {
		t.Fatalf("page after reopen: %v", e)
	}
	if len(page) != 2 {
		t.Fatalf("after reopen retained %d versions, want 2", len(page))
	}
	if page[0].Version != 1 || page[0].Content != "第一版" {
		t.Fatalf("after reopen page[0] = %+v", page[0])
	}
	if page[1].Version != 2 || page[1].Content != "第二版" {
		t.Fatalf("after reopen page[1] = %+v", page[1])
	}

	// A new save after reopening must continue the version sequence rather
	// than restart at 1 (which would overwrite version 1).
	a3, e := svc2.Save("doc-1", "owner", "第三版")
	if e != nil {
		t.Fatalf("save 3 after reopen: %v", e)
	}
	if a3.Version != 3 {
		t.Fatalf("save 3 after reopen: version = %d, want 3", a3.Version)
	}
	page, _, _ = svc2.Page("doc-1", 1, 100)
	if len(page) != 3 || page[2].Content != "第三版" {
		t.Fatalf("third version not retained: %+v", page)
	}
	if page[0].Content != "第一版" {
		t.Fatalf("first version lost after reopen+save: %+v", page[0])
	}
}

// TestSavePaginates exercises the "page through all annotations" look-back
// workflow step and confirms paging is stable across the saved versions.
func TestSavePaginates(t *testing.T) {
	s, _ := openStore(t, "page.db")
	defer s.Close()
	svc := New(s)

	const total = 7
	for i := 0; i < total; i++ {
		if _, e := svc.Save("doc-1", "owner", "c"+itoa(i+1)); e != nil {
			t.Fatalf("save %d: %v", i+1, e)
		}
	}

	var collected []domain.Annotation
	pageNo, size := 1, 3
	for {
		page, more, e := svc.Page("doc-1", pageNo, size)
		if e != nil {
			t.Fatalf("page %d: %v", pageNo, e)
		}
		collected = append(collected, page...)
		if !more {
			break
		}
		pageNo++
	}
	if len(collected) != total {
		t.Fatalf("paged through %d, want %d", len(collected), total)
	}
	for i, a := range collected {
		if a.Version != i+1 {
			t.Fatalf("collected[%d].Version = %d, want %d", i, a.Version, i+1)
		}
	}
}

func itoa(n int) string {
	if n == 0 {
		return "0"
	}
	var b []byte
	for n > 0 {
		b = append([]byte{byte('0' + n%10)}, b...)
		n /= 10
	}
	return string(b)
}
