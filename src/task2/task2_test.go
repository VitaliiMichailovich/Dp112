package task2

import "testing"

var testCase = []struct {
	e1   Envelope
	e2   Envelope
	want int
	werr string
}{
	{
		Envelope{15.0, 20.0},
		Envelope{14.0, 18.0},
		2,
		"",
	},
	{
		Envelope{13.0, 20.0},
		Envelope{14.0, 18.0},
		0,
		"",
	},
	{
		Envelope{1.0, 1.0},
		Envelope{2.0, 2.0},
		1,
		"",
	},
	{
		Envelope{1.1, 1.1},
		Envelope{1.2, 1.2},
		1,
		"",
	},
	{
		Envelope{2.9, 2.9},
		Envelope{3.0, 3.0},
		1,
		"",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 77.78},
		1,
		"",
	},
	{
		Envelope{1, 77.77},
		Envelope{77.78, 77.78},
		1,
		"",
	},
	{
		Envelope{0, 77.77},
		Envelope{77.78, 77.78},
		0,
		"Incorrect EnvelopeA size. Every size of envelope \"{0 77.77}\" must be > 0.",
	},
	{
		Envelope{77.77, 0},
		Envelope{77.78, 77.78},
		0,
		"Incorrect EnvelopeA size. Every size of envelope \"{77.77 0}\" must be > 0.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{0, 77.78},
		0,
		"Incorrect EnvelopeB size. Every size of envelope \"{0 77.78}\" must be > 0.",
	},
	{
		Envelope{77.77, 77.77},
		Envelope{77.78, 0},
		0,
		"Incorrect EnvelopeB size. Every size of envelope \"{77.78 0}\" must be > 0.",
	},
}

func TestTask2(t *testing.T) {
	for _, tc := range testCase {
		got, err := Task2(tc.e1, tc.e2)
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

func BenchmarkTask2(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			Task2(tc.e1, tc.e2)
		}
	}
}
