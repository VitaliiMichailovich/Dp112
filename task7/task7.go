package task7

import (
	"errors"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

type Params struct {
	Context string `json:"context"`
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

func doTask7(path string) ([]int64, error) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return []int64{}, fmt.Errorf("Problem with opening 'context' file.\n%v", err)
	}
	defer file.Close()
	var text = make([]byte, 128)
	for {
		n, err := file.Read(text)
		if err != nil && err != io.EOF {
			return []int64{}, fmt.Errorf("Problem with opening 'context' file.\n%v", err)
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
	if err != nil {
		return []int64{}, fmt.Errorf("Incorrect input. Expect int.\n%v", err)
	}
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
		if err != nil {
			return []int64{}, errors.New("Incorrect input. Expect int and int.")
		}
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
	return fib[start:stop], nil
}

func task7validator(path string) ([]int64, error) {
	var file, err = os.OpenFile(path, os.O_RDWR, 0644)
	if err != nil {
		return []int64{}, errors.New("Problem with opening 'context' file.")
	}
	var text = make([]byte, 128)
	for {
		n, err := file.Read(text)
		if err != nil && err != io.EOF {
			return []int64{}, errors.New("Problem with reading 'context' file.")
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
	_, err = strconv.Atoi(inI[0])
	if err != nil {
		return []int64{}, errors.New("Incorrect input. Expect int.")
	}
	if len(inI) == 2 {
		_, err = strconv.Atoi(inI[1])
		if err != nil {
			return []int64{}, errors.New("Incorrect input. Expect int and int.")
		}
	}
	file.Close()
	task7Result, err := doTask7(path)
	if err != nil {
		return task7Result, err
	}
	return task7Result, nil
}

func Task(param Params) ([]int64, error) {
	return task7validator(param.Context)
}