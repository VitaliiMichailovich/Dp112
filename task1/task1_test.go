package task1

import (
	"encoding/json"
	"errors"
	"testing"
)

var testCaseTask = []struct {
	w    int
	h    int
	s    string
	want string
	werr error
}{
	{
		1,
		1,
		"a",
		"a",
		nil,
	},
	{
		2,
		2,
		"@",
		"@ \n @",
		nil,
	},
	{
		5,
		3,
		"%",
		"% % %\n % % \n% % %",
		nil,
	},
	{
		8,
		8,
		"#",
		"# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #\n# # # # \n # # # #",
		nil,
	},
	{
		8,
		-8,
		"*",
		"",
		errors.New("Incorrect input. Width (8) and height (-8) must be greater than 0 and symbol (*) cant be nil."),
	},
	{
		-8,
		8,
		"Z",
		"",
		errors.New("Incorrect input. Width (-8) and height (8) must be greater than 0 and symbol (Z) cant be nil."),
	},
	{
		8,
		8,
		"7",
		"7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7\n7 7 7 7 \n 7 7 7 7",
		nil,
	},
	{
		0,
		0,
		"o",
		"",
		errors.New("Incorrect input. Width (0) and height (0) must be greater than 0 and symbol (o) cant be nil."),
	},
	{
		7,
		77,
		"",
		"",
		errors.New("Incorrect input. Width (7) and height (77) must be greater than 0 and symbol () cant be nil."),
	},
	{
		5,
		1,
		"@",
		"@ @ @",
		nil,
	},
}

func TestTask(t *testing.T) {
	for _, tc := range testCaseTask {
		param, _ := json.Marshal(Params{
			Width:  tc.w,
			Height: tc.h,
			Symbol: tc.s,
		})
		got, err := Task(param)
		if err != nil && err.Error() != tc.werr.Error() {
			t.Fatalf("Task1(%v, %v, %q) got error \n%v\nwant \n%v",
				tc.h, tc.w, tc.s, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task1(%v, %v, %q) = %q, want %q",
				tc.h, tc.w, tc.s, got, tc.want)
		}
	}
}

func BenchmarkTask(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for _, tc := range testCaseTask {
			param, _ := json.Marshal(Params{
				Width:  tc.w,
				Height: tc.h,
				Symbol: tc.s,
			})
			Task(param)
		}
	}
}
