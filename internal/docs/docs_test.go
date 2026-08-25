package docs

import "testing"

func TestAnnotationVersionsDistinct(t *testing.T) {
	a := NewGuide("a", "one")
	b := Revise(a, "two")
	if !Distinct(a, b) {
		t.Fatal("annotation revisions must remain distinct")
	}
}
func TestGuideValidity(t *testing.T) {
	if !Valid(NewGuide("a", "b")) {
		t.Fatal("invalid guide")
	}
}
