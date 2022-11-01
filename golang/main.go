package main

import (
	"math/rand"
	"sort"
	"time"
)

// tuneable parameters
// TODO: extract these to somewhere nicer
const numPoints int = 256  // setting this over 63 might cause problems
const maxSize float64 = 10 // basically the dimensions of the world
const maxPopulation int = 1000
const numGenerations int = 50
const numSurvivors int = 100

func main() {

	rand.Seed(time.Now().UnixNano())

	// determine the points of interest in our little world
	var points = make([]Point, numPoints)
	for i := range points {
		points[i] = newRandomPoint(i, maxSize)
	}

	// setup the first generation
	var population = make([]Agent, maxPopulation)
	for i := range population {
		population[i] = newRandomAgent(numPoints)
	}

	// TODO: keep track of best score for each generation
	for i := 0; i < numGenerations; i++ {
		evaluatePopulation(population, points)

		// get the slice of the surviving agents
		var survivors = population[:numSurvivors]

		// create the next generation
		for i := range population {
			population[i] = newAgentFromParent(survivors)
			population[i].mutate()
		}
	}

	// evaluate the final population
	evaluatePopulation(population, points)
}

// helper function to evaluate and sort a population
func evaluatePopulation(population []Agent, points []Point) {
	for i := range population {
		population[i].evaluate(points)
	}
	sort.Sort(ByScore(population))
}
