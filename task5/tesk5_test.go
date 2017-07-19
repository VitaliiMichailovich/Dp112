package task5

import (
	"encoding/json"
	"testing"
)

var testCase = []struct {
	min  int
	max  int
	want string
	werr string
}{
	{
		75319,
		100952,
		"easy, 1672, 838",
		"",
	},
	{
		905319,
		909952,
		"easy, 312, 149",
		"",
	},
	{
		0,
		999999,
		"easy, 55252, 25081",
		"",
	},
	{
		-2,
		999999,
		", 0, 0",
		"Incorrect input. Minimal ticket number (-2) must be >= 0 and Maximal (999999) ticket must be >0. ",
	},
	{
		7,
		-5,
		", 0, 0",
		"Incorrect input. Minimal ticket number (7) must be >= 0 and Maximal (-5) ticket must be >0. ",
	},
	{
		-8,
		-2,
		", 0, 0",
		"Incorrect input. Minimal ticket number (-8) must be >= 0 and Maximal (-2) ticket must be >0. ",
	},
	{
		1999999,
		2999999,
		", 0, 0",
		"Incorrect input. To big tickets (1999999, 2999999). Ticket vust be < 999999.",
	},
	{
		0,
		1999999,
		", 0, 0",
		"Incorrect input. To big tickets (0, 1999999). Ticket vust be < 999999.",
	},
	{
		9999999,
		5,
		", 0, 0",
		"Incorrect input. Minimal ticket number (9999999) must be <= Maximal (5) ticket. ",
	},
}

func TestTask5(t *testing.T) {
	for _, tc := range testCase {
		in, _ := json.Marshal(Params{Min: tc.min, Max: tc.max})
		got, err := Task(in)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task5(%d, %d) got error \n%v\nwant \n%v",
				tc.min, tc.max, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task5(%d, %d) = %v, want %v",
				tc.min, tc.max, got, tc.want)
		}
	}
}

func BenchmarkTask5(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			in, _ := json.Marshal(Params{Min: tc.min, Max: tc.max})
			Task(in)
		}
	}
}
