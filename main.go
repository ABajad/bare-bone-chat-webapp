package main

import (
	"log"
	"net/http"
	"path/filepath"
	"text/template"
)

type templ struct {
	source string
	templ  *template.Template
}

func (t *templ) Handle(w http.ResponseWriter, r *http.Request) {
	if t.templ == nil {
		t.templ = template.Must(template.ParseFiles(filepath.Join("templates", t.source)))
	}
	t.templ.Execute(w, nil)
}

func main() {
	r := newRoom()
	http.HandleFunc("/", (&templ{source: "chat.html"}).Handle)
	http.Handle("/room", r)
	go r.run()
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal("ListenAndServe:", err)
	}
}
