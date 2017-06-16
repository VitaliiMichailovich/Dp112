//package task3
package main

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
			return out, errors.New("\""+tri.name+"\" is not a Triangle")
		}
		squares = append(squares, square{tri.name, s})
	}

	sort.Slice(squares, func(i, j int) bool {return squares[i].sq < squares[j].sq})
	for _, a := range squares {
		out = append(out, a.name)
	}
	return out, nil
}

func main() {
	fmt.Println(Task3(Triangle{"j1", 10, 20, 25}, Triangle{"j2", 60, 22.5, 77.8}, Triangle{"j3", 22.2, 20.5, 30.8}, Triangle{"j4", 1.1, 2.71, 3.4}, Triangle{"j5", 0.18, 0.20, 0.30}, Triangle{"j6", 270, 270, 3.0}, Triangle{"j7", 1750, 1750.20, 10.30}))
}
