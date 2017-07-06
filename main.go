package main

import (
	"github.com/VitaliiMichailovich/DP112/src/server"
)
const path = "./tasks.json"

func main() {
	server.Server()

/*
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
	fmt.Println(task1.Task(params.Params1))
	fmt.Println("********** Task 2: **********")
	fmt.Println(task2.Task(params.Params2))
	fmt.Println("********** Task 3: **********")
	fmt.Println(task3.Task(params.Params3))
	fmt.Println("********** Task 4: **********")
	fmt.Println(task4.Task(params.Params4))
	fmt.Println("********** Task 5: **********")
	fmt.Println(task5.Task(params.Params5))
	fmt.Println("********** Task 6: **********")
	fmt.Println(task6.Task(params.Params6))
	fmt.Println("********** Task 7: **********")
	fmt.Println(task7.Task(params.Params7))
*/
}
