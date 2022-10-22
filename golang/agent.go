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

// ByScore implements sort.Interface for []Agent based on the Agent.score field
type ByScore []Agent

func (a ByScore) Len() int           { return len(a) }
func (a ByScore) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByScore) Less(i, j int) bool { return a[i].score < a[j].score }

// a new agent with random genes
func newRandomAgent(geneSize int) Agent {

	// initialise genes
	var genes = make([]int, geneSize)
	for i := range genes {
		genes[i] = i
	}

	// shuffle
	rand.Shuffle(len(genes), func(i, j int) { genes[i], genes[j] = genes[j], genes[i] })

	return Agent{genes, math.Inf(1)}
}

// a new empty agent
func newAgent(geneSize int) Agent {
	return Agent{make([]int, geneSize), math.Inf(1)}
}

// new agent which is a clone of a random parent
func newAgentFromParent(parents []Agent) Agent {
	// TODO: make this more interesting

	// choose a random parent
	parent := parents[rand.Intn(len(parents))]

	// clone the parent
	var newAgent = newAgent(len(parent.genes))
	copy(newAgent.genes, parent.genes)

	return newAgent
}

// a function to mutate an agent by swapping two points
func (a *Agent) mutate() {
	// TODO: make this more interesting

	// choose two random genes
	i, j := rand.Intn(len(a.genes)), rand.Intn(len(a.genes))

	// swap the genes
	a.genes[i], a.genes[j] = a.genes[j], a.genes[i]
}

// a function to evaluate the fitness of an agent
func (a *Agent) evaluate(points []Point) (score float64) {

	// traverse the path and add the distance to each point
	for i := 0; i < len(a.genes)-1; i++ {

		score += distance(points[a.genes[i]], points[a.genes[i+1]])
	}

	// add the distance back to the start
	score += distance(points[a.genes[len(a.genes)-1]], points[a.genes[0]])

	a.score = score
	return
}
