package main

import (
	"fmt"
	"math"
	"math/rand"
)

type Point struct {
	idx  int // not quite sure why this is here...
	X, Y float64
}

func distance(p, q Point) float64 {
	return math.Hypot(p.X-q.X, p.Y-q.Y)
}

func newRandomPoint(index int, lim float64) Point {
	return Point{
		idx: index,
		X:   rand.Float64() * lim,
		Y:   rand.Float64() * lim,
	}
}

func (p Point) String() string {
	return fmt.Sprintf("%2d: (%.3f, %.3f)", p.idx, p.X, p.Y)
}
