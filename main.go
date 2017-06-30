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
	"server"
)
const path = "./tasks.json"

type ParamsData struct {
	Chess		ChessParams		`json:"task1"`
	Envelopes	[]task2.Envelope	`json:"task2"`
	Triangles	[]task3.Triangle	`json:"task3"`
	Palindrom	int64			`json:"task4"`
	Tickets		[]int			`json:"task5"`
	Sequence	[]int			`json:"task6"`
	Fibonacci	string			`json:"task7"`
}

type ChessParams struct {
	Width	int	`json:"width"`
	Height	int	`json:"height"`
	Symbol	string	`json:"symbol"`
}

func main() {
	server.Server()
	params := ParamsData{}
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Println("Can't read input file.")
	}
	err = json.Unmarshal(raw, &params)
	if err != nil {
		fmt.Println("Can't parse json.")
	}
	fmt.Println("********** Task1: **********")
	fmt.Println(task1.Task1(params.Chess.Width, params.Chess.Height, params.Chess.Symbol))
	fmt.Println("********** Task2: **********")
	fmt.Println(task2.Task2(params.Envelopes[0], params.Envelopes[1]))
	fmt.Println("********** Task3: **********")
	fmt.Println(task3.Task3(params.Triangles))
	fmt.Println("********** Task4: **********")
	fmt.Println(task4.Task4(params.Palindrom))
	fmt.Println("********** Task5: **********")
	fmt.Println(task5.Task5(params.Tickets[0], params.Tickets[1]))
	fmt.Println("********** Task6: **********")
	fmt.Println(task6.Task6(params.Sequence[0], params.Sequence[1]))
	fmt.Println("********** Task7: **********")
	fmt.Println(task7.Task7())
}
