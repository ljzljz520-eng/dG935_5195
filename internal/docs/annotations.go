package docs

type AnnotationGuide struct {
	Title, Body string
	Revision    int
}

func NewGuide(title, body string) AnnotationGuide {
	return AnnotationGuide{Title: title, Body: body, Revision: 1}
}
func Revise(g AnnotationGuide, body string) AnnotationGuide { g.Body = body; g.Revision++; return g }
func Distinct(a, b AnnotationGuide) bool                    { return a.Revision != b.Revision || a.Body != b.Body }
func Valid(g AnnotationGuide) bool                          { return g.Title != "" && g.Body != "" && g.Revision > 0 }
