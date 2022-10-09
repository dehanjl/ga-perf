package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct {
	idx  int
	X, Y float64
}

// distance bewteen two points
func distance(p, q Point) float64 {
	return math.Hypot(float64(p.X-q.X), float64(p.Y-q.Y))
}

func newPoint(index int, lim float64) Point {
	return Point{
		idx: index,
		X:   rand.Float64() * lim,
		Y:   rand.Float64() * lim,
	}
}

func (p Point) String() string {
	return fmt.Sprintf("%2d: (%.3f, %.3f)", p.idx, p.X, p.Y)
}
