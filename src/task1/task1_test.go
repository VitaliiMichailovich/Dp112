package task1

import (
	"testing"
)

var testCase = []struct {
	h    int
	w    int
	s    string
	want string
	werr string
}{
	{
		1,
		1,
		"a",
		"a",
		"",
	},
	{
		2,
		2,
		"@",
		"@ \n @",
		"",
	},
	{
		5,
		3,
		"%",
		"% % %\n % % \n% % %",
		"",
	},
	{
		8,
		8,
		"#",
		"# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #",
		"",
	},
	{
		8,
		-8,
		"*",
		"",
		"Incorrect input. Width (8) and height (-8) must be greater than 0 and symbol (*) cant be nil.",
	},
	{
		-8,
		8,
		"Z",
		"",
		"Incorrect input. Width (-8) and height (8) must be greater than 0 and symbol (Z) cant be nil.",
	},
	{
		8,
		8,
		"7",
		"7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7",
		"",
	},
	{
		0,
		0,
		"o",
		"",
		"Incorrect input. Width (0) and height (0) must be greater than 0 and symbol (o) cant be nil.",
	},
	{
		7,
		77,
		"",
		"",
		"Incorrect input. Width (7) and height (77) must be greater than 0 and symbol () cant be nil.",
	},

	{
		9999999,
		1,
		"@",
		"",
		"",
	},
}

func TestTask1(t *testing.T) {
	for _, tc := range testCase {
		got, err := task1Validator(tc.h, tc.w, tc.s)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task1(%d, %d, %q) got error \n%v\nwant \n%v",
				tc.h, tc.w, tc.s, err, tc.werr)
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
			task1Validator(tc.h, tc.w, tc.s)
		}
	}
}
