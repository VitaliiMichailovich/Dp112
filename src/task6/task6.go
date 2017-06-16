package task6

import (
	"fmt"
	"math"
	"os"
	"strconv"
)

const path = "./task6.txt"

func Task6(length, pow int) {
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
			fmt.Println(err.Error())
			os.Exit(0)
		}
	}
	var file, err = os.Create(path)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	defer file.Close()
	_, err = file.WriteString(toFile)
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
	err = file.Sync()
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
