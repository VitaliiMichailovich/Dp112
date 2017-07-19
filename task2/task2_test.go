package task2

import (
	"encoding/json"
	"errors"
	"testing"
)

var testCaseTask = []struct {
	e1   Envelope
	e2   Envelope
	want string
	werr error
}{
	{
		Envelope{15.0, 20.0},
		Envelope{14.0, 18.0},
		"2",
		nil,
	},
	{
		Envelope{13.0, 20.0},
		Envelope{14.0, 18.0},
		"0",
		nil,
	},
	{
		Envelope{1.0, 1.0},
		Envelope{2.0, 2.0},
		"1",
		nil,
	},
	{
		Envelope{1.1, 1.1},
		Envelope{1.2, 1.2},
		"1",
		nil,
	},
	{
		Envelope{2.9, 2.9},
		Envelope{3.0, 3.0},
		"1",
		nil,
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 77.78},
		"1",
		nil,
	},
	{
		Envelope{1, 77.77},
		Envelope{77.78, 77.78},
		"1",
		nil,
	},
	{
		Envelope{0, 77.77},
		Envelope{77.78, 77.78},
		"0",
		errors.New("Incorrect EnvelopeA size. Every size of envelope \"{0 77.77}\" must be > 0."),
	},
	{
		Envelope{77.77, 0},
		Envelope{77.78, 77.78},
		"0",
		errors.New("Incorrect EnvelopeA size. Every size of envelope \"{77.77 0}\" must be > 0."),
	},
	{
		Envelope{-77.77, -77.77},
		Envelope{77.78, 77.78},
		"0",
		errors.New("Incorrect EnvelopeA size. Every size of envelope \"{-77.77 -77.77}\" must be > 0."),
	},
	{
		Envelope{77.77, 77.77},
		Envelope{0, 77.78},
		"0",
		errors.New("Incorrect EnvelopeB size. Every size of envelope \"{0 77.78}\" must be > 0."),
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 0},
		"0",
		errors.New("Incorrect EnvelopeB size. Every size of envelope \"{77.78 0}\" must be > 0."),
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, -5},
		"0",
		errors.New("Incorrect EnvelopeB size. Every size of envelope \"{77.78 -5}\" must be > 0."),
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.77, 77.77},
		"0",
		errors.New("Some of envelope sizes are equal. Check Input parametres. env1: {77.77 77.77}; env2: {77.77 77.77}"),
	},
}

func TestTask(t *testing.T) {
	for _, tc := range testCaseTask {
		param, _ := json.Marshal(Params{
			Envelope1: tc.e1,
			Envelope2: tc.e2,
		})
		got, err := Task(param)
		if err != nil && err.Error() != tc.werr.Error() {
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
