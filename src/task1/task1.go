package task1

import (
	"fmt"
	"strconv"
)

type Params struct {
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Symbol string `json:"symbol"`
}

// Task "Chess board" if all inputs is correct.
func doTask1(w int, h int, s string) string {
	var out string
	for i := 1; i <= h; i++ {
		for ii := 1; ii <= w; ii++ {
			if (ii+i)%2 != 0 {
				out += " "
			} else {
				out += s
			}
		}
		if i != h {
			out += "\n"
		}
	}
	return out
}

// Pre validator for task "Chess board"
func task1Validator(widthInterface, heightInterface, symbolInterface interface{}) (string, error) {
	// Check width
	width, ok := widthInterface.(int)
	if !ok {
		widthS, err := strconv.Atoi(widthInterface.(string))
		if err != nil {
			return "", fmt.Errorf("Incorrect input of width: \"%s\". Must be int32.", widthInterface)
		}
		width = widthS
	}
	// Check height
	height, ok := heightInterface.(int)
	if !ok {
		heightS, err := strconv.Atoi(heightInterface.(string))
		if err != nil {
			return "", fmt.Errorf("Incorrect input of height: \"%s\". Must be int32.", heightInterface)
		}
		height = heightS
	}
	// Check string
	symbol, ok := symbolInterface.(string)
	if !ok {
		return "", fmt.Errorf("Incorrect input of height: \"%s\". Must be int32.", symbolInterface)
	}
	// Check correct
	if width <= 0 || height <= 0 || symbol == "" {
		return "", fmt.Errorf("Incorrect input. Width (%d) and height (%d) must be greater than 0 and symbol (%s) cant be nil.", width, height, symbol)
	}
	return doTask1(width, height, symbol), nil
}

func Task(param Params) (string, error) {
	return task1Validator(param.Width, param.Height, param.Symbol)
}