package task7

import (
	"testing"
)

var testCase = []struct {
	want []int64
	werr string
}{
	/*
	{
		[]int64{1597, 2584, 4181, 6765},
		"",
	},
	*/
	{
		[]int64{},
		"Problem with opening 'context' file.",
	},
}

func TestTask7(t *testing.T) {
	for _, tc := range testCase {
		got, err := Task7()
		if err != nil && err.Error() != tc.werr{
			t.Fatalf("Task7() got error \n%v\nwant \n%v",
				err, tc.werr)
		}
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
