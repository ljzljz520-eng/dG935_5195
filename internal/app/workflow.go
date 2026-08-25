package app

import (
	"accountingcollab/internal/annotations"
	"accountingcollab/internal/domain"
	ex "accountingcollab/internal/export"
	"accountingcollab/internal/review"
	"accountingcollab/internal/store"
)

type App struct {
	Store       *store.Store
	Annotations *annotations.Service
	Reviews     *review.Service
}

func New(s *store.Store) *App {
	return &App{Store: s, Annotations: annotations.New(s), Reviews: review.New(s)}
}
func (a *App) StartWorkspace(id, name, owner string) (domain.Workspace, error) {
	w := domain.NewWorkspace(id, name, owner)
	return w, a.Store.PutWorkspace(w)
}
func (a *App) AddDocument(id, wid, title string) (domain.Document, error) {
	d := domain.NewDocument(id, wid, title)
	return d, a.Store.PutDocument(d)
}
func (a *App) AddAnnotation(doc, author, content string) (domain.Annotation, error) {
	return a.Annotations.Save(doc, author, content)
}
func (a *App) CompleteReview(doc, reviewer, decision string) (domain.ReviewTask, error) {
	r, e := a.Reviews.Open(doc, reviewer)
	if e != nil {
		return r, e
	}
	return a.Reviews.Complete(r.ID, decision)
}
func (a *App) Export(doc string) (ex.Snapshot, error) { return ex.Build(a.Store, doc) }
func (a *App) Close() error                           { return a.Store.Close() }
