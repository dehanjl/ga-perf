package main

// tuneable parameters
// TODO: extract these to somewhere nicer
const numPoints int = 10
const maxSize float64 = 10
const maxPopulation int = 20
const numGenerations int = 20

func main() {

	// determine the points of interest in our little world
	var points = make([]Point, numPoints)
	for i := range points {
		points[i] = newPoint(i, maxSize)
	}

	// setup the first generation
	var population = make([]Agent, maxPopulation)
	for i := range population {
		population[i] = newRandomAgent(numPoints * 2)
	}

}
