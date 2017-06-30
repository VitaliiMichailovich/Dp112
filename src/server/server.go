package server

import (
	"net/http"
	"path/filepath"
	"text/template"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("static", "index.html")
	fp := filepath.Join("static", filepath.Clean(r.URL.Path))
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func HandleTask1(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()

	w.Write([]byte("<h1>Task1</h1>"))
}

func HandleTask2(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte("<h1>Task2</h1>"))
}

func Server() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	//http.HandleFunc("/", IndexPage)
	http.HandleFunc("/task/1", HandleTask1)
	http.HandleFunc("/task/2", HandleTask2)
	http.ListenAndServe(":8088", nil)
}
