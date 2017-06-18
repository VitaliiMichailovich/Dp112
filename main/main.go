package main

import (
	"../src/task1"
	//"../src/task2"
	//"../src/task3"
	//"../src/task4"
	//"../src/task5"
	//"../src/task6"
	//"../src/task7"
	"encoding/json"
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
	"errors"
)

const path = "./tasks.json"

func main() {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	var text = make([]byte, 1024)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	var textCorr []byte
	for k, v := range text {
		if v == 0 {
			textCorr = text[0:k]
			break
		}
	}
	var dat map[string]map[string]string
	if err := json.Unmarshal(textCorr, &dat); err != nil {
		panic(err)
	}
	for k, v := range dat {
		for kTask, vTask := range v {
			switch k {
			case "task1":
				w, h, s, err := task1Validator(vTask)
				if err == nil {
					fmt.Println("Task1, " + kTask + ": \n" + task1.Task1(w, h, s) + "\n")
				} else {
					checkError(err)
				}
			}
		}
	}
}

func task1Validator(in string) (int, int, string, error) {
	inW := strings.Split(in, ",")
	w, err := strconv.Atoi(inW[0])
	if err != nil {
		return 0, 0, "", errors.New("Incorrect input. Expect int, int, symbol and coma as a separator.")
	}
	h, err := strconv.Atoi(inW[1])
	if err != nil {
		return 0, 0, "", errors.New("Incorrect input. Expect int, int, symbol and coma as a separator.")
	}
	s := inW[2]
	return w, h, s, nil
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
