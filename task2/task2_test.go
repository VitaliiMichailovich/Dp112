package task2

import (
	"encoding/json"
	"testing"
)

var testCaseTask = []struct {
	e1   Envelope
	e2   Envelope
	want string
	werr string
}{
	{
		Envelope{15.0, 20.0},
		Envelope{14.0, 18.0},
		"2",
		"",
	},
	{
		Envelope{13.0, 20.0},
		Envelope{14.0, 18.0},
		"0",
		"",
	},
	{
		Envelope{1.0, 1.0},
		Envelope{2.0, 2.0},
		"1",
		"",
	},
	{
		Envelope{1.1, 1.1},
		Envelope{1.2, 1.2},
		"1",
		"",
	},
	{
		Envelope{2.9, 2.9},
		Envelope{3.0, 3.0},
		"1",
		"",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 77.78},
		"1",
		"",
	},
	{
		Envelope{1, 77.77},
		Envelope{77.78, 77.78},
		"1",
		"",
	},
	{
		Envelope{0, 77.77},
		Envelope{77.78, 77.78},
		"0",
		"Incorrect EnvelopeA size. Every size of envelope \"{0 77.77}\" must be greater than zero.",
	},
	{
		Envelope{77.77, 0},
		Envelope{77.78, 77.78},
		"0",
		"Incorrect EnvelopeA size. Every size of envelope \"{77.77 0}\" must be greater than zero.",
	},
	{
		Envelope{-77.77, -77.77},
		Envelope{77.78, 77.78},
		"0",
		"Incorrect EnvelopeA size. Every size of envelope \"{-77.77 -77.77}\" must be greater than zero.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{0, 77.78},
		"0",
		"Incorrect EnvelopeB size. Every size of envelope \"{0 77.78}\" must be greater than zero.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 0},
		"0",
		"Incorrect EnvelopeB size. Every size of envelope \"{77.78 0}\" must be greater than zero.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, -5},
		"0",
		"Incorrect EnvelopeB size. Every size of envelope \"{77.78 -5}\" must be greater than zero.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.77, 77.77},
		"0",
		"Some of envelope sizes are equal. Check Input parametres. env1: {77.77 77.77}; env2: {77.77 77.77}",
	},
}

func TestTask(t *testing.T) {
	for _, tc := range testCaseTask {
		param, _ := json.Marshal([2]Envelope{
			{AB: tc.e1.AB, CD: tc.e1.CD},
			{AB: tc.e2.AB, CD: tc.e2.CD},
		})
		got, err := Task(param)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task2(%f, %f) got error \n%v\nwant \n%v",
				tc.e1, tc.e2, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task2(%f, %f) = %v, want %d",
				tc.e1, tc.e2, got, tc.want)
		}
	}
}

func BenchmarkTask(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCaseTask {
			param, _ := json.Marshal(Params{
				Envelope1: tc.e1,
				Envelope2: tc.e2,
			})
			Task(param)
		}
	}
}
