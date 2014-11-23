package main

import (
	"net/http"
	"fmt"
	"io/ioutil"
)

func main() {
	http.HandleFunc("/", MyHandler)
	http.ListenAndServe(":4747", nil)
}

func MyHandler(rw http.ResponseWriter, req *http.Request) {
	page := loadPage("index")
	fmt.Fprintln(rw, string(page.Body))
}

type Page struct {
	Title string
	Body []byte
}

func loadPage(title string) *Page {
	filename := title + ".html"
	body, _ := ioutil.ReadFile(filename)
	return &Page{Title: title, Body: body}
}