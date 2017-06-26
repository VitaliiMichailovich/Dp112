package task1

import (
	"fmt"
	"strconv"
)

func DoTask1(w int, h int, s string) string {
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

func Task1(widthInterface, heightInterface, symbolInterface interface{}) (string, error) {
	width, ok := widthInterface.(int)
	if ok == false {
		widthS, err := strconv.Atoi(widthInterface.(string))
		if err != nil {
			return "", fmt.Errorf("Incorrect input of width: \"%s\". Must be int32.", widthInterface)
		}
		width = widthS
	}
	height, ok := heightInterface.(int)
	if ok == false {
		heightS, err := strconv.Atoi(heightInterface.(string))
		if err != nil {
			return "", fmt.Errorf("Incorrect input of height: \"%s\". Must be int32.", heightInterface)
		}
		height = heightS
	}
	symbol, ok := symbolInterface.(string)
	if ok == false {
		return "", fmt.Errorf("Incorrect input of height: \"%s\". Must be int32.", symbolInterface)
	}
	if width <= 0 || height <= 0 || symbol == "" {
		return "", fmt.Errorf("Incorrect input. Width (%d) and height (%d) must be greater than 0 and symbol (%s) cant be nil.", width, height, symbol)
	}
	return DoTask1(width, height, symbol), nil
}