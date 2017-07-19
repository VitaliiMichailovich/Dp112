package taskregister

import (
	"errors"
	"fmt"
)

var TaskRegister map[int]func([]byte) (string, error) = make(map[int]func([]byte) (string, error))

func InitializeTask(task int, runnerFunc func([]byte) (string, error)) {
	TaskRegister[task] = runnerFunc
}

func RunTask(task int, jsonParams []byte) (string, error) {
	taskFunc, ok := TaskRegister[task]
	if !ok {
		return "", errors.New(fmt.Sprintf("Task #%v didn't registered runner", task))
	}
	return taskFunc(jsonParams)
}
