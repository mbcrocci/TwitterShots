package main

import (
	"io/ioutil"

	"github.com/go-martini/martini"
)

func main() {
	m := martini.Classic()
	page := loadPage("index")

	m.Get("/", func() string {
		return string(page.Body)
	})
	m.Run()
}

type Page struct {
	Title string
	Body  []byte
}

func loadPage(title string) *Page {
	filename := title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}
