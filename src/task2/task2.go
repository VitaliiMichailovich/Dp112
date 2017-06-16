package task2

type Envelope struct {
	ab, cd float32
}

func Task2(envA, envB Envelope) int {
	out := 0
	if envA.ab <= 0 || envA.cd <= 0 || envB.ab <= 0 || envB.cd <= 0 {
		return out
	}
	if envA.cd > envA.ab {
		envA.ab, envA.cd = envA.cd, envA.ab
	}
	if envB.cd > envB.ab {
		envB.ab, envB.cd = envB.cd, envB.ab
	}
	if (envA.ab > envB.ab) && (envA.cd > envB.cd) {
		out = 2
	}
	if (envB.ab > envA.ab) && (envB.cd > envA.cd) {
		out = 1
	}
	return out
}
