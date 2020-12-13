package view

import (
	"GoPractice/crawler/frontend/model"
	template2 "html/template"
	"os"
	"testing"
)

func TestTemplate(t *testing.T) {
	template := template2.Must(
		template2.ParseFiles("template.html"))

	page := model.SearchResult{}
	err := template.Execute(os.Stdout, page)
	if err != nil {
		panic(err)
	}
}
