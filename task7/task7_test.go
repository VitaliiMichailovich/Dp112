package task7

import (
	"encoding/json"
	"testing"
)

var testCase = []struct {
	in   string
	want string
	werr string
}{
	{
		"context",
		"1597, 2584, 4181, 6765",
		"",
	},
	{
		"context1",
		"10946, 17711, 28657, 46368, 75025",
		"",
	},
	{
		"context2",
		"",
		"Problem with opening 'context' file.",
	},
}

func TestTask7(t *testing.T) {
	for _, tc := range testCase {
		in, _ := json.Marshal(Params{Context: tc.in})
		got, err := Task(in)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task7(%v) got error \n%v\nwant \n%v",
				tc.in, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task7(%v) = %s, want %s",
				tc.in, got, tc.want)
		}
	}
}

func BenchmarkTask7(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			in, _ := json.Marshal(Params{Context: tc.in})
			Task(in)
		}
	}
}
