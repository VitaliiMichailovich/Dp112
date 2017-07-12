package server

import (
	"encoding/json"
	"fmt"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"io/ioutil"
	"net/http"
	"strconv"
	"html/template"
)

type WriteBack struct {
	Resp, Reason string
}

//func IndexPage(w http.ResponseWriter, r *http.Request) {
//	lp := filepath.Join("static", "index.html")
//	fp := filepath.Join("static", filepath.Clean(r.URL.Path))
//	tmpl, _ := template.ParseFiles(lp, fp)
//	tmpl.ExecuteTemplate(w, "layout", nil)
//}

func HandleTask(w http.ResponseWriter, r *http.Request) {
	var params map[string]json.RawMessage
	//taskIdInURI := strings.Split(r.RequestURI, "/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	err = json.Unmarshal(body, &params)
	//var taskId int
	//for _, val := range taskIdInURI {
	//	taskId, err = strconv.Atoi(val)
	//	if err == nil{
	//		break
	//	}
	//}
	var value []byte
	var task int
	for t, v := range params {
		value = v
		task, err = strconv.Atoi(t[4:])
		if err != nil {
			fmt.Sprintf("Sorry, but I can't find a task number in file (%v).", t)
		}
	}
	result, err := taskregister.RunTask(task, value)
	errString := ""
	if err != nil {
		errString = string(err.Error())
	}
	back, _ := json.Marshal(WriteBack{
		Resp:   result,
		Reason: errString,
	})
	w.Header().Set("Content-Type", "application/json")
	w.Write(back)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	var params map[string]json.RawMessage
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	err = json.Unmarshal(body, &params)
	var write []WriteBack
	for t, v := range params {
		task, err := strconv.Atoi(t[4:])
		if err != nil {
			fmt.Sprintf("Sorry, but I can't find a task number in file (%v).", t)
		}
		result, err := taskregister.RunTask(task, v)
		errString := ""
		if err != nil {
			errString = string(err.Error())
		}
		write = append(write, WriteBack{result,errString})
	}
	back, _ := json.Marshal(write)
	w.Header().Set("Content-Type", "application/json")
	w.Write(back)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("client/app/templates/tasks.html")
	if err != nil {
		fmt.Fprintf(w, err.Error())
		return
	}

	t.ExecuteTemplate(w, "index", nil)
}

func Server() {
	http.Handle("/client/", http.StripPrefix("/client/", http.FileServer(http.Dir("./client/"))))
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/tasks", HandleTasks)
	http.HandleFunc("/task/", HandleTask)
	http.ListenAndServe(":8089", nil)
}
