package task4

import (
	"testing"
	"encoding/json"
)

var testCase = []struct {
	in   int64
	want string
	werr string
}{
	{
		123457468437,
		"0",
		"There is NO Palindrome. ",
	},
	{
		778891198877,
		"778891198877",
		"",
	},
	{
		7788911988771,
		"778891198877",
		"",
	},
	{
		778891198877,
		"778891198877",
		"",
	},
	{
		17788887787888878,
		"87888878",
		"",
	},
	{
		18788887877888877,
		"87888878",
		"",
	},
	{
		99991999999999999,
		"999999999999",
		"",
	},
	{
		5,
		"0",
		"Incorrect input \"5\". Input must be greater than 10.",
	},
}

func TestTask4(t *testing.T) {
	for _, tc := range testCase {
		in, _ := json.Marshal(Params{Number:tc.in})
		got, err := Task(in)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task4(%v) got error \n%v\nwant \n%v",
				tc.in, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task4(%d) = %d, want %d",
				tc.in, got, tc.want)
		}
	}
}

func BenchmarkTask4(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			in, _ := json.Marshal(Params{Number:tc.in})
			Task(in)
		}
	}
}
