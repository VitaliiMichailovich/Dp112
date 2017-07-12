package server

import (
	"encoding/json"
	"fmt"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"html/template"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
)

type WriteBack struct {
	Task   int    `json:"task"`
	Resp   string `json:"resp"`
	Reason string `json:"reason"`
}

//func IndexPage(w http.ResponseWriter, r *http.Request) {
//	lp := filepath.Join("static", "index.html")
//	fp := filepath.Join("static", filepath.Clean(r.URL.Path))
//	tmpl, _ := template.ParseFiles(lp, fp)
//	tmpl.ExecuteTemplate(w, "layout", nil)
//}

func HandleTask(w http.ResponseWriter, r *http.Request) {
	//var params map[string]json.RawMessage
	taskIdInURI := strings.Split(r.RequestURI, "/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	//err = json.Unmarshal(body, &params)
	var taskId int
	for _, val := range taskIdInURI {
		taskId, err = strconv.Atoi(val)
		if err == nil {
			break
		} else if len(val) > 4 {
			taskId, err = strconv.Atoi(val[4:])
			if err == nil {
				break
			}
		}
	}
	if taskId == 0 {
		back, _ := json.Marshal(WriteBack{
			Task:   0,
			Resp:   "",
			Reason: "I did not find any number of task in link.",
		})
		w.Header().Set("Content-Type", "application/json")
		w.Write(back)
	}
	//var value []byte
	//var task int
	//for t, v := range params {
	//	value = v
	//	fmt.Println(string(value), string(params))
	//	task, err = strconv.Atoi(t[4:])
	//	if err != nil {
	//		fmt.Sprintf("Sorry, but I can't find a task number in file (%v).", t)
	//	}
	//}
	result, err := taskregister.RunTask(taskId, body)
	errString := ""
	if err != nil {
		errString = string(err.Error())
	}
	back, _ := json.Marshal(WriteBack{
		Task: taskId,
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
		write = append(write, WriteBack{task, result, errString})
	}
	back, _ := json.Marshal(write)
	w.Header().Set("Content-Type", "application/json")
	w.Write(back)
}

func IndexHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("templates/index.html")
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
