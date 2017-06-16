//package task7
package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

const path = "./context"

func Task7() []int {
	var out []int

	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	defer file.Close()

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
	fib := []int64{0, 1}
	for i := 2; i < 93; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}

	in := strings.Fields(string(text))
	fmt.Println(in[1])
	if len(in) == 1 {
		//len, err := strconv.Atoi(in[0])

	} else {
		var start, stop int
		min, err := strconv.ParseInt(in[0], 10, 64)
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(0)
		}
		max, err := strconv.ParseInt(in[1], 10, 64)
		if err != nil{
			fmt.Println(err.Error())
			os.Exit(0)
		}
		for i := 0; i < len(fib); i++ {
			if fib[i] >= min && start == 0 {
				start = i
			}
			if fib[i] >= max {
				stop = i-1
			}
		}
		fmt.Println(start, stop)
	}
	return out
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}

func main() {
	fmt.Println(Task7())
}
