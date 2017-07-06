package server

import (
	"net/http"
	"path/filepath"
	"text/template"
	"github.com/VitaliiMichailovich/DP112/src/task1"
	"github.com/VitaliiMichailovich/DP112/src/task2"
	"github.com/VitaliiMichailovich/DP112/src/task3"
	"github.com/VitaliiMichailovich/DP112/src/task4"
	"github.com/VitaliiMichailovich/DP112/src/task5"
	"github.com/VitaliiMichailovich/DP112/src/task6"
	"github.com/VitaliiMichailovich/DP112/src/task7"
	"encoding/json"
	"fmt"
	"strconv"
)

type Params struct {
	Params1	task1.Params	`json:"task1"`
	Params2	task2.Params	`json:"task2"`
	Params3	task3.Params	`json:"task3"`
	Params4	task4.Params	`json:"task4"`
	Params5	task5.Params	`json:"task5"`
	Params6	task6.Params	`json:"task6"`
	Params7	task7.Params	`json:"task7"`
}

func IndexPage(w http.ResponseWriter, r *http.Request) {
	lp := filepath.Join("static", "index.html")
	fp := filepath.Join("static", filepath.Clean(r.URL.Path))
	tmpl, _ := template.ParseFiles(lp, fp)
	tmpl.ExecuteTemplate(w, "layout", nil)
}

func HandleTask1(w http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	t1 := Params{}
	err = decoder.Decode(&t1)
	if err != nil {
		fmt.Println(err.Error())
	}
	resp1, err := task1.Task(t1.Params1)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(resp1))
}

func HandleTask2(w http.ResponseWriter, r *http.Request) {
	var err error
	defer r.Body.Close()
	decoder := json.NewDecoder(r.Body)
	t2 := Params{}
	err = decoder.Decode(&t2)
	if err != nil {
		fmt.Println(err.Error())
	}
	defer r.Body.Close()
	resp2, err := task2.Task(t2.Params2)
	if err != nil {
		w.Write([]byte(err.Error()))
	}
	w.Write([]byte(strconv.Itoa(resp2)))
}

func Server() {
	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)

	http.HandleFunc("/task/1", HandleTask1)
	http.HandleFunc("/task/2", HandleTask2)
	http.ListenAndServe(":8089", nil)
}
