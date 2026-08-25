package domain

import "time"

type Workspace struct {
	ID, Name, Owner string
	CreatedAt       time.Time
}
type Document struct {
	ID, WorkspaceID, Title, Status string
	UpdatedAt                      time.Time
}
type Annotation struct {
	ID, DocumentID, Author, Content string
	Version                         int
	CreatedAt                       time.Time
}
type ReviewTask struct {
	ID, DocumentID, Reviewer, Decision string
	CompletedAt                        time.Time
}

func NewWorkspace(id, name, owner string) Workspace {
	return Workspace{ID: id, Name: name, Owner: owner, CreatedAt: time.Now().UTC()}
}
func NewDocument(id, workspace, title string) Document {
	return Document{ID: id, WorkspaceID: workspace, Title: title, Status: "open", UpdatedAt: time.Now().UTC()}
}
func NewAnnotation(id, doc, author, content string, version int) Annotation {
	return Annotation{ID: id, DocumentID: doc, Author: author, Content: content, Version: version, CreatedAt: time.Now().UTC()}
}
func NewReviewTask(id, doc, reviewer string) ReviewTask {
	return ReviewTask{ID: id, DocumentID: doc, Reviewer: reviewer, Decision: "pending"}
}
func ValidateWorkspace(w Workspace) error {
	if w.ID == "" || w.Name == "" || w.Owner == "" {
		return ErrInvalid
	}
	return nil
}
func ValidateDocument(d Document) error {
	if d.ID == "" || d.WorkspaceID == "" || d.Title == "" {
		return ErrInvalid
	}
	return nil
}
func ValidateAnnotation(a Annotation) error {
	if a.DocumentID == "" || a.Author == "" || a.Content == "" || a.Version < 1 {
		return ErrInvalid
	}
	return nil
}
func ValidateReview(r ReviewTask) error {
	if r.ID == "" || r.DocumentID == "" || r.Reviewer == "" {
		return ErrInvalid
	}
	return nil
}
