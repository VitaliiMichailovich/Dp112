package task6

import (
	"testing"
	"encoding/json"
)

var testCase = []struct {
	len  int
	pow  int
	werr string
}{
	{
		7,
		777,
		"",
	},
	{
		7,
		7777,
		"",
	},
	{
		-2,
		77,
		"Incorrect input: \"-2,77\". Must be > 0.",
	},
	{
		2,
		-77,
		"Incorrect input: \"2,-77\". Must be > 0.",
	},
	{
		0,
		-77,
		"Incorrect input: \"0,-77\". Must be > 0.",
	},
}

func TestTask6(t *testing.T) {
	for _, tc := range testCase {
		in, _ := json.Marshal(Params{Length:tc.len, Pow:tc.pow})
		_, err := Task(in)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task6(%d, %d) got error \n%v\nwant \n%v",
				tc.len, tc.pow, err, tc.werr)
		}
	}
}

func BenchmarkTask6(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			in, _ := json.Marshal(Params{Length:tc.len, Pow:tc.pow})
			Task(in)
		}
	}
}
