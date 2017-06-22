package task2

type Envelope struct {
	AB, CD float32
}

func Task2(envA, envB Envelope) int {
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
