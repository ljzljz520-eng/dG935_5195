package domain

import "testing"

func TestValidationRules(t *testing.T) {
	if ValidateWorkspace(Workspace{}) == nil || ValidateDocument(Document{}) == nil || ValidateAnnotation(Annotation{}) == nil {
		t.Fatal()
	}
	if !CanTransition("open", "review") || CanTransition("archived", "open") {
		t.Fatal()
	}
}
