package task7

import (
	"fmt"
	"io"
	"os"
	"strings"
	"strconv"
)

const path = "./context"

func Task7() []int64 {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	checkError(err)
	var text = make([]byte, 128)
	for {
		n, err := file.Read(text)
		if err != io.EOF {
			checkError(err)
		}
		if n == 0 {
			break
		}
	}
	var inS string
	for k, v := range text {
		if v == 0 {
			inS = string(text[0:k])
			break
		}
	}
	inI := strings.Split(inS, " ")

	fib := []int64{0, 1}
	for i := 2; i < 93; i++ {
		fib = append(fib, fib[i-1]+fib[i-2])
	}
	var min, max int
	min, err = strconv.Atoi(inI[0])
	checkError(err)

	var start, stop int
	if len(inI) == 1 {
		for k, v := range fib {
			bit := getBitNumber(v)
			if bit == int64(min) && start == 0 {
				start = k
			} else if bit > int64(min) {
				stop = k
				break
			}
		}
	} else {
		max, err = strconv.Atoi(inI[1])
		checkError(err)
		for k, v := range fib {
			if v >= int64(min) && start == 0 {
				start = k
			}
			if v >= int64(max) {
				stop = k
			}
			if stop > start {
				break
			}
		}
	}
	return fib[start:stop]
}

func getBitNumber(in int64) int64 {
	var out int64
	for {
		in = in / 10
		out++
		if in == 0 {
			break
		}
	}
	return out
}

func checkError(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}