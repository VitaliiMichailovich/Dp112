package task3

import (
	"fmt"
	"math"
	"sort"
)

type Triangle struct {
	Name string	`json: "name"`
	A    float32	`json: "a"`
	B    float32	`json: "b"`
	C    float32	`json: "c"`
}

type Triangler interface {
	Triangle()
}

type square struct {
	name string
	sq   float64
}

func doTask3(triangles []Triangle) []string {
	var out []string
	var squares []square
	for _, tri := range triangles {
		p := (tri.A + tri.B + tri.C) / 2
		s := math.Sqrt(float64(p * (p - tri.A) * (p - tri.B) * (p - tri.C)))
		squares = append(squares, square{tri.Name, s})
	}
	sort.Slice(squares, func(i, j int) bool { return squares[i].sq < squares[j].sq })
	for _, a := range squares {
		out = append(out, a.name)
	}
	return out
}

func Task3(in []Triangle) ([]string, error) {
	var triangles []Triangle
	for _, triangleTMP := range in {
		if triangleTMP.A <= 0 || triangleTMP.B <= 0 || triangleTMP.C <= 0 {
			return []string{""}, fmt.Errorf("Incorrect input of triangle: \"%v\". All sides of triangle must be > 0", triangleTMP)
		}
		if len(triangleTMP.Name) != 3 {
			return []string{""}, fmt.Errorf("Incorrect input of triangle name \"%v\". Must be Three letters.", triangleTMP.Name)
		}
		var a, b, c float32
		if triangleTMP.A >= triangleTMP.B && triangleTMP.A >= triangleTMP.C {
			a = triangleTMP.A
			b = triangleTMP.B
			c = triangleTMP.C
		} else if triangleTMP.B >= triangleTMP.A && triangleTMP.B >= triangleTMP.C {
			a = triangleTMP.B
			b = triangleTMP.A
			c = triangleTMP.C
		} else {
			a = triangleTMP.C
			b = triangleTMP.A
			c = triangleTMP.B
		}
		if (b + c) <= a {
			return []string{""}, fmt.Errorf("This is not a triangle (\"%v\"). Specify the incoming data.", triangleTMP)
		}
		triangles = append(triangles, triangleTMP)
	}
	return doTask3(triangles), nil
}
