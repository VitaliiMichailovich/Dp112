package task3

import (
	"encoding/json"
	"testing"
)

var testCase = []struct {
	t1   []Triangle
	want string
	werr string
}{
	{
		[]Triangle{{"abc", 18, 20, 22}, {"bcd", 60, 62.5, 77.8}},
		"bcd,abc",
		"",
	},
	{
		[]Triangle{{"abc", 10, 20, 22}, {"bcd", 60, 22.5, 77.8}, {"def", 22.2, 20.5, 30.8}, {"efg", 1.1, 2.71, 3.4}, {"fgh", 0.18, 0.20, 0.30}, {"ghi", 270, 270, 3.0}, {"hij", 1750, 1750.20, 10.30}},
		"hij,bcd,ghi,def,abc,efg,fgh",
		"",
	},
	{
		[]Triangle{{"abc", 100, 20, 22}, {"bcd", 60, 22.5, 77.8}},
		"",
		"This is not a triangle (\"{abc 100 20 22}\"). Specify the incoming data.",
	},
	{
		[]Triangle{{"abc", 10, 20, 22}, {"bcd", 600, 22.5, 77.8}},
		"",
		"This is not a triangle (\"{bcd 600 22.5 77.8}\"). Specify the incoming data.",
	},
	{
		[]Triangle{{"abc", 0, 20, 22}, {"bcd", 600, 22.5, 77.8}},
		"",
		"Incorrect input of triangle: \"{abc 0 20 22}\". All sides of triangle must be > 0",
	},
	{
		[]Triangle{{"abc", 10, 0, 22}, {"bcd", 600, 22.5, 77.8}},
		"",
		"Incorrect input of triangle: \"{abc 10 0 22}\". All sides of triangle must be > 0",
	},
	{
		[]Triangle{{"abc", 10, 20, 0}, {"bcd", 600, 22.5, 77.8}},
		"",
		"Incorrect input of triangle: \"{abc 10 20 0}\". All sides of triangle must be > 0",
	},
	{
		[]Triangle{{"abc", 10, 20, 22}, {"bcd", 00, 22.5, 77.8}},
		"",
		"Incorrect input of triangle: \"{bcd 0 22.5 77.8}\". All sides of triangle must be > 0",
	},
	{
		[]Triangle{{"abc", 10, 20, 22}, {"bcd", 10, 0, 77.8}},
		"",
		"Incorrect input of triangle: \"{bcd 10 0 77.8}\". All sides of triangle must be > 0",
	},
	{
		[]Triangle{{"abc", 10, 20, 22}, {"bcd", 10, 22.5, 0}},
		"",
		"Incorrect input of triangle: \"{bcd 10 22.5 0}\". All sides of triangle must be > 0",
	},
}

func TestTask3(t *testing.T) {
	for _, tc := range testCase {
		var param []Triangle
		for _, v := range tc.t1 {
			param = append(param, Triangle{Name:v.Name, A:v.A, B:v.B,C:v.C},)
		}
		paramTask, _ := json.Marshal(param)
		got, err := Task(paramTask)
		if err != nil && err.Error() != tc.werr {
			t.Fatalf("Task3(%v) got error \n%v\nwant \n%v",
				tc.t1, err, tc.werr)
		}
		if got != tc.want {
			t.Fatalf("Task3(%v) = %s, want %s",
				tc.t1, got, tc.want)
		}

	}
}

func BenchmarkTask3(b *testing.B) {
	for i := 0; i <= b.N; i++ {
		for _, tc := range testCase {
			param, _ := json.Marshal(Params{
				tc.t1,
			})
			Task(param)
		}
	}
}
