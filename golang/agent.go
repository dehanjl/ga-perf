package main

import (
	"math"
	"math/rand"
)

// agent struct which contains the order in which it visits the points
type Agent struct {
	genes []int
	score float64
}

func newRandomAgent(geneSize int) Agent {

	var genes = make([]int, geneSize)
	for i := range genes {
		genes[i] = rand.Intn(numPoints)
	}
	return Agent{genes, math.Inf(-1)}
}
