package export

import (
	"accountingcollab/internal/domain"
	"accountingcollab/internal/store"
	"fmt"
	"strings"
)

type Snapshot struct {
	Document    domain.Document
	Annotations []domain.Annotation
	Summary     string
}

func Build(s *store.Store, docID string) (Snapshot, error) {
	d, e := s.GetDocument(docID)
	if e != nil {
		return Snapshot{}, e
	}
	a, e := s.ListAnnotations(docID)
	if e != nil {
		return Snapshot{}, e
	}
	return Snapshot{Document: d, Annotations: a, Summary: fmt.Sprintf("%s:%d", d.Title, len(a))}, nil
}
func Render(v Snapshot) string {
	parts := []string{v.Document.Title, v.Document.Status}
	for _, a := range v.Annotations {
		parts = append(parts, fmt.Sprintf("v%d %s", a.Version, a.Content))
	}
	return strings.Join(parts, " | ")
}
func FilterByAuthor(in []domain.Annotation, author string) []domain.Annotation {
	out := []domain.Annotation{}
	for _, a := range in {
		if a.Author == author {
			out = append(out, a)
		}
	}
	return out
}
