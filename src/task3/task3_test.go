package task3

import (
	"testing"
)

var testCase = []struct {
	t1   []Triangle
	want []string
}{
	{
		[]Triangle{{"j1", 18, 20, 725}, {"j2", 60, 62.5, 77.8}},
		[]string{"j1", "j2"},
	},
	{
		[]Triangle{{"j1", 10, 20, 725}, {"j2", 60, 22.5, 77.8}, {"j3", 22.2, 20.5, 30.8}, {"j4", 1.1, 2.71, 3.4}, {"j5", 0.18, 0.20, 0.30}, {"j6", 270, 270, 3.0}, {"j7", 1750, 1750.20, 10.30}},
		[]string{"j1", "j5", "j4", "j3", "j6", "j2", "j7"},
	},
}

func TestTask3(t *testing.T) {
	for _, tc := range testCase {
		got := Task3(tc.t1)
		for k, gotId := range got {
			if gotId != tc.want[k] {
				t.Fatalf("Task3(%s) = %s, want %s",
					tc.t1, got, tc.want)
			}
		}
	}
}

func BenchmarkTask3(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			Task3(tc.t1)
		}
	}
}
