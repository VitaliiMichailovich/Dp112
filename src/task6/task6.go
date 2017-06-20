package task6

import (
	"math"
	"os"
	"strconv"
	"errors"
)

const path = "./task6.txt"

func Task6(length, pow int) error {
	var toFile string
	start := math.Sqrt(float64(pow))
	if float64(int(start)) < start {
		start = float64(int(start)) + 1
	}
	for i := 0; i < length; i++ {
		toFile += strconv.Itoa(int(start) + i)
		if i != length-1 {
			toFile += ", "
		}
	}
	if _, err := os.Stat(path); os.IsExist(err) {
		err := os.Remove(path)
		if err != nil {
			return errors.New(err.Error())
		}
	}
	var file, err = os.Create(path)
	if err != nil {
		return errors.New(err.Error())
	}
	defer file.Close()
	_, err = file.WriteString(toFile)
	if err != nil {
		return errors.New(err.Error())
	}
	err = file.Sync()
	if err != nil {
		return errors.New(err.Error())
	}
	return nil
}