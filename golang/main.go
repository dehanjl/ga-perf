package main

import (
	"fmt"
	"sort"
	"time"
)

// tuneable parameters
// TODO: extract these to somewhere nicer
const numPoints int = 32   // setting this over 63 might cause problems
const maxSize float64 = 10 // basically the dimensions of the world
const maxPopulation int = 1000
const numGenerations int = 50
const numSurvivors int = 100

func main() {

	// determine the points of interest in our little world
	var points = make([]Point, numPoints)
	for i := range points {
		points[i] = newRandomPoint(i, maxSize)
	}

	startTime := time.Now()

	// setup the first generation
	var population = make([]Agent, maxPopulation)
	for i := range population {
		population[i] = newRandomAgent(numPoints * 2)
	}

	for i := 0; i < numGenerations; i++ {
		fmt.Println("Generation", i+1)

		// evaluate the population
		evaluatePopulation(population, points)

		fmt.Println("Best score:", population[0].score)

		// get the slice of the surviving agents
		var survivors = population[:numSurvivors]

		// create the next generation
		for i := range population {
			population[i] = newAgentFromParents(survivors)
			population[i].mutate()
		}
	}

	fmt.Println("Generation", numGenerations+1)

	evaluatePopulation(population, points)

	fmt.Println("Best score:", population[0].score)

	elapsed := time.Since(startTime)
	fmt.Println("Execution time:", elapsed)
	fmt.Println("Time per generation:", elapsed/time.Duration(numGenerations+1))
}

// helper function to evaluate and sort a population
func evaluatePopulation(population []Agent, points []Point) {
	for i := range population {
		population[i].evaluate(points)
	}
	sort.Sort(ByScore(population))
}
