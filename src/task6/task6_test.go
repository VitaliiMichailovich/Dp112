package task6

import "testing"

var testCase = []struct{
	len int
	pow int
	want error
}{
	{
		7,
		777,
		nil,
	},
	{
		7,
		7777,
		nil,
	},
}

func TestTask6(t *testing.T) {
	for _, tc := range testCase {
		got := Task6(tc.len, tc.pow)
		if got != tc.want {
			t.Fatalf("Task6(%d, %d) = %s, want %s",
				tc.len, tc.pow, got, tc.want)
		}
	}
}

func BenchmarkTask6(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			Task6(tc.len, tc.pow)
		}
	}
}