//package task7
package main

import (
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

const path = "./context"

func Task7() []int64 {
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

	var start, stop int
	if len(in) == 1 {
		//len, err := strconv.Atoi(in[0])
	} else {
		min, err := strconv.ParseInt(in[0], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		max, err := strconv.ParseInt(in[1], 10, 64)
		if err != nil {
			fmt.Println(err.Error())
			os.Exit(0)
		}
		for k, v := range fib {
			if v >= min && start == 0 {
				start = k
			}
			if v >= max {
				stop = k
			}
			fmt.Println(k, v, start, stop, min, max)
			if stop > start {
				break
			}
		}
	}
	return fib[start:stop]
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
