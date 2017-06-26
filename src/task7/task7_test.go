package task7

import (
	"testing"
)

var testCase = []struct {
	want []int64
}{
	{
		[]int64{1597, 2584, 4181, 6765},
	},
}

func TestTask7(t *testing.T) {
	for _, tc := range testCase {
		got := Task7()
		for gotId, gotVal := range got {
			if gotVal != tc.want[gotId] {
				t.Fatalf("Task7() = %s, want %s",
					got, tc.want)
			}
		}
	}
}

func BenchmarkTask7(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		Task7()
	}
}
