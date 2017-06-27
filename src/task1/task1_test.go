package task1

import (
	"testing"
)

var testCase = []struct {
	h    int
	w    int
	s    string
	want string
}{
	{
		1,
		1,
		"a",
		"a",
	},
	{
		2,
		2,
		"@",
		"@ \n @",
	},
	{
		5,
		3,
		"%",
		"% % %\n % % \n% % %",
	},
	{
		8,
		8,
		"#",
		"# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #",
	},
	/*
	{
		8,
		-8,
		"*",
		"Incorrect input. Width (8) and height (-8) must be greater than 0 and symbol (*) cant be nil.",
	},
	*/
}

func TestTask1(t *testing.T) {
	for _, tc := range testCase {
		got, err := Task1(tc.h, tc.w, tc.s)
		if err != nil && err != tc.want{
			t.Fatalf("Task1(%d, %d, %q) = %v, want %q",
				tc.h, tc.w, tc.s, err, tc.want)
		}
		if got != tc.want {
			t.Fatalf("Task1(%d, %d, %q) = %q, want %q",
				tc.h, tc.w, tc.s, got, tc.want)
		}
	}
}

func BenchmarkTask1(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCase {
			Task1(tc.h, tc.w, tc.s)
		}
	}
}
