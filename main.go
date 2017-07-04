package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"task1"
	"task2"
	"task3"
	"task4"
	"task5"
	"task6"
	"task7"
	//"server"
)
const path = "./tasks.json"

type Params struct {
	Params1	[]task1.Params	`json:"task1"`
	Params2	[]task2.Params	`json:"task2"`
	Params3	[]task3.Params	`json:"task3"`
	Params4	[]task4.Params	`json:"task4"`
	Params5	[]task5.Params	`json:"task5"`
	Params6	[]task6.Params	`json:"task6"`
	Params7	[]task7.Params	`json:"task7"`
}

func main() {
	//server.Server()
	params := Params{}
	raw, err := ioutil.ReadFile(path)

	if err != nil {
		fmt.Println("Can't read input file.")
	}
	err = json.Unmarshal(raw, &params)
	if err != nil {
		fmt.Println("Can't parse json.")
	}
	fmt.Println("********** Task 1: **********")
	for _, param := range params.Params1 {
		fmt.Println(task1.Task(param))
	}
	fmt.Println("********** Task 2: **********")
	for _, param := range params.Params2 {
		fmt.Println(task2.Task(param))
	}
	fmt.Println("********** Task 3: **********")
	for _, param := range params.Params3 {
		fmt.Println(task3.Task(param))
	}
	fmt.Println("********** Task 4: **********")
	for _, param := range params.Params4 {
		fmt.Println(task4.Task(param))
	}
	fmt.Println("********** Task 5: **********")
	for _, param := range params.Params5 {
		fmt.Println(task5.Task(param))
	}
	fmt.Println("********** Task 6: **********")
	for _, param := range params.Params6 {
		fmt.Println(task6.Task(param))
	}
	fmt.Println("********** Task 7: **********")
	for _, param := range params.Params7 {
		fmt.Println(task7.Task(param))
	}
}
