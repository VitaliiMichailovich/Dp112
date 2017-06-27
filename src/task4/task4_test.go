package task4

import "testing"

var testCase = []struct {
	in   int64
	want int64
}{
	{
		123457468437,
		0,
	},
	{
		778891198877,
		778891198877,
	},
	{
		7788911988771,
		778891198877,
	},
	{
		17788887787888878,
		87888878,
	},
	{
		18788887877888877,
		87888878,
	},
}

func TestTask4(t *testing.T) {
	for _, tc := range testCase {
		got, _ := Task4(tc.in)
		if got != tc.want {
			t.Fatalf("Task4(%d) = %d, want %d",
				tc.in, got, tc.want)
		}
	}
}

func BenchmarkTask4(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			Task4(tc.in)
		}
	}
}
