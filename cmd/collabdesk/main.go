package main

import (
	"accountingcollab/internal/app"
	"accountingcollab/internal/store"
	"fmt"
	"os"
)

func main() {
	path := "collabdesk.db"
	if len(os.Args) > 1 {
		path = os.Args[1]
	}
	s, e := store.Open(path)
	if e != nil {
		panic(e)
	}
	defer s.Close()
	a := app.New(s)
	w, e := a.StartWorkspace("default", "会计资料协作盘", "owner")
	if e != nil {
		panic(e)
	}
	d, e := a.AddDocument("doc-1", w.ID, "季度凭证")
	if e != nil {
		panic(e)
	}
	_, _ = a.AddAnnotation(d.ID, "owner", "请核对金额")
	fmt.Println(d.Title)
}
