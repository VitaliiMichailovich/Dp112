package server

import (
	"net/http"
	"strings"
	"io/ioutil"
	"strconv"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"encoding/json"
)

type WriteBack struct{
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
	taskIdInURI := strings.Split(r.RequestURI, "/")
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}
	err = json.Unmarshal(body, &params)
	var taskId int
	for _, val := range taskIdInURI {
		taskId, err = strconv.Atoi(val)
		if err == nil{
			break
		}
	}
	var value []byte
	for _, v := range params {
		value = v
	}
	result, err := taskregister.RunTask(taskId, value)
	errString := ""
	if err != nil {
		errString = string(err.Error())
	}

	back, _ := json.Marshal(WriteBack{
		Resp: result,
		Reason: errString,
	})
	w.Write(back)
}

func HandleTasks(w http.ResponseWriter, r *http.Request) {
	//var err error
	//defer r.Body.Close()
	//decoder := json.NewDecoder(r.Body)
	//t2 := Params{}
	//err = decoder.Decode(&t2)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//defer r.Body.Close()
	//resp2, err := task2.Task(t2.Params2)
	//if err != nil {
	//	w.Write([]byte(err.Error()))
	//}
	w.Write([]byte("BBB"))
}

func Server() {

	fileServer := http.FileServer(http.Dir("static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/tasks", HandleTasks)
	http.HandleFunc("/task/", HandleTask)
	http.ListenAndServe(":8089", nil)
}
