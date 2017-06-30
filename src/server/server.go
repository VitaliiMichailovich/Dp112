package server

import (
	"net/http"
)

func IndexPage(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	w.Write([]byte("<a href=\"task/1\">Task1</a><br><a href=\"task/2\">Task2</a>"))
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
	http.HandleFunc("/", IndexPage)
	http.HandleFunc("/task/1", HandleTask1)
	http.HandleFunc("/task/2", HandleTask2)
	http.ListenAndServe(":8088", nil)
}
