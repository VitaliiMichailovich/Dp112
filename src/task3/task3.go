package task3

import (
	"errors"
	"math"
	"sort"
)

type Triangle struct {
	Name string
	A    float32
	B    float32
	C    float32
}

type square struct {
	name string
	sq   float64
}

func Task3(triangles []Triangle) ([]string, error) {
	var out []string
	var squares []square
	for _, tri := range triangles {
		p := (tri.A + tri.B + tri.C) / 2
		s := math.Sqrt(float64(p * (p - tri.A) * (p - tri.B) * (p - tri.C)))
		if math.IsNaN(s) {
			return out, errors.New("\"" + tri.Name + "\" is not a Triangle")
		}
		squares = append(squares, square{tri.Name, s})
	}

	sort.Slice(squares, func(i, j int) bool { return squares[i].sq < squares[j].sq })
	for _, a := range squares {
		out = append(out, a.name)
	}
	return out, nil
}