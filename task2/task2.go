package task2

import (
	"encoding/json"
	"fmt"
	"github.com/VitaliiMichailovich/DP112/taskregister"
	"strconv"
)

type Params struct {
	Envelope1 Envelope `json:"envelope1"`
	Envelope2 Envelope `json:"envelope2"`
}

type Envelope struct {
	AB float32 `json:"width"`
	CD float32 `json:"height"`
}

type Enveloper interface {
	Envelope()
}

func Task(bytesParams []byte) (string, error) {
	var param [2]Envelope
	err := json.Unmarshal([]byte(bytesParams), &param)
	if err != nil {
		return "", err
	}
	ret, err := task2validator(Envelope{param[0].AB, param[0].CD}, Envelope{param[1].AB, param[1].CD})
	return strconv.Itoa(ret), err
}


func init() {
	taskregister.InitializeTask(2, Task)
}

// Task "Envelope analysis" if all inputs is correct.
func doTask2(envA, envB Envelope) int {
	out := 0
	if (envA.AB > envB.AB) && (envA.CD > envB.CD) {
		out = 2
	}
	if (envB.AB > envA.AB) && (envB.CD > envA.CD) {
		out = 1
	}
	return out
}

// Task "Envelope analysis" pre validator. If all OK: run doTask2
func task2validator(envA, envB interface{}) (int, error) {
	// check envelope 1
	env1, ok := envA.(Envelope)
	if !ok {
		return 0, fmt.Errorf("Incorrect input of envA: \"%v\". Must be \"Envelope\".", envA)
	}
	if env1.AB <= 0 || env1.CD <= 0 {
		return 0, fmt.Errorf("Incorrect EnvelopeA size. Every size of envelope \"%v\" must be greater than zero.", env1)
	}
	env2, ok := envB.(Envelope)
	if !ok {
		return 0, fmt.Errorf("Incorrect input of envB: \"%v\". Must be \"Envelope\".", envB)
	}
	if env2.AB <= 0 || env2.CD <= 0 {
		return 0, fmt.Errorf("Incorrect EnvelopeB size. Every size of envelope \"%v\" must be greater than zero.", env2)
	}
	if env1.CD > env1.AB {
		env1.AB, env1.CD = env1.CD, env1.AB
	}
	if env2.CD > env2.AB {
		env2.AB, env2.CD = env2.CD, env2.AB
	}
	if (env1.AB == env2.AB) || (env1.CD == env2.CD) {
		return 0, fmt.Errorf("Some of envelope sizes are equal. Check Input parametres. env1: %v; env2: %v", envA, envB)
	}
	return doTask2(env1, env2), nil
}