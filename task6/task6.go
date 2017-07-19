package task6

import (
	"encoding/json"
	"fmt"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"math"
	"os"
	"strconv"
)

type Params struct {
	Length int `json:"length"`
	Pow    int `json:"pow"`
}

const path = "task6.txt"

func Task(bytesParams []byte) (string, error) {
	var param Params
	err := json.Unmarshal(bytesParams, &param)
	if err != nil {
		return "", err
	}
	return "", task6Validator(param.Length, param.Pow)
}

func init() {
	taskregister.InitializeTask(6, Task)
}

func doTask6(length, pow int) error {
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
			return err
		}
	}
	var file, err = os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(toFile)
	if err != nil {
		return err
	}
	err = file.Sync()
	if err != nil {
		return err
	}
	return nil
}

func task6Validator(length, pow interface{}) error {
	lengthChecked, ok := length.(int)
	if !ok {
		lengthI, err := strconv.Atoi(length.(string))
		if err != nil {
			return fmt.Errorf("Incorrect input: \"%v\". Must be int32. \n%v", length, err)
		}
		lengthChecked = lengthI
	}
	powChecked, ok := pow.(int)
	if !ok {
		powI, err := strconv.Atoi(pow.(string))
		if err != nil {
			return fmt.Errorf("Incorrect input: \"%v\". Must be int32. \n%v", pow, err)
		}
		powChecked = powI
	}
	if lengthChecked < 0 || powChecked < 0 {
		return fmt.Errorf("Incorrect input: \"%v,%v\". Must be > 0.", lengthChecked, powChecked)
	}
	return doTask6(lengthChecked, powChecked)
}
