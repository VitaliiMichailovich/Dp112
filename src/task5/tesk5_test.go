package task5

import "testing"

var testCase = []struct{
	min int
	max int
	want Ticket
} {
	{
		75319,
		100952,
		Ticket{"easy", 1672, 810},
	},
	{
		905319,
		909952,
		Ticket{"easy", 312, 35},
	},
}

func TestTask5(t *testing.T) {
	for _, tc := range testCase {
		got := Task5(tc.min, tc.max)
		if got != tc.want {
			t.Fatalf("Task5(%d, %d) = %d, want %d",
				tc.min, tc.max, got, tc.want)
		}
	}
}

func BenchmarkTask5(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			Task5(tc.min, tc.max)
		}
	}
}