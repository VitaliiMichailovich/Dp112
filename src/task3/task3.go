package task3

import (
	"errors"
	"fmt"
	"math"
	"sort"
)

type Triangle struct {
	name string
	a    float32
	b    float32
	c    float32
}

type square struct {
	name string
	sq   float64
}

func Task3(triangles ...Triangle) ([]string, error) {
	var out []string
	var squares []square
	for _, tri := range triangles {
		p := (tri.a + tri.b + tri.c) / 2
		s := math.Sqrt(float64(p * (p - tri.a) * (p - tri.b) * (p - tri.c)))
		if math.IsNaN(s) {
			return out, errors.New("\"" + tri.name + "\" is not a Triangle")
		}
		squares = append(squares, square{tri.name, s})
	}

	sort.Slice(squares, func(i, j int) bool { return squares[i].sq < squares[j].sq })
	for _, a := range squares {
		out = append(out, a.name)
	}
	return out, nil
}
