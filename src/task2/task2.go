package task2

import "fmt"

type Params struct {
	Envelope1 Envelope `json:"envelope1"`
	Envelope2 Envelope `json:"envelope2"`
}

type Envelope struct {
	AB float32 `json:"ab"`
	CD float32 `json:"cd"`
}

type Enveloper interface {
	Envelope()
}

// Task "Envelope analysis" if all inputs is correct.
func doTask2(envA, envB Envelope) int {
	out := 0
	if envA.CD > envA.AB {
		envA.AB, envA.CD = envA.CD, envA.AB
	}
	if envB.CD > envB.AB {
		envB.AB, envB.CD = envB.CD, envB.AB
	}
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
		return 0, fmt.Errorf("Incorrect EnvelopeA size. Every size of envelope \"%v\" must be > 0.", env1)
	}
	// check envelope 1
	env2, ok := envB.(Envelope)
	if !ok {
		return 0, fmt.Errorf("Incorrect input of envB: \"%v\". Must be \"Envelope\".", envB)
	}
	if env2.AB <= 0 || env2.CD <= 0 {
		return 0, fmt.Errorf("Incorrect EnvelopeB size. Every size of envelope \"%v\" must be > 0.", env2)
	}
	return doTask2(env1, env2), nil
}

func Task (param Params) (int, error){
	return task2validator(param.Envelope1, param.Envelope2)
}