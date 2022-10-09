package main

const numPoints int = 10
const maxSize float64 = 10
const maxPopulation int = 20
const numGenerations int = 20

func main() {

	var points = make([]Point, numPoints)
	for i := range points {
		points[i] = newPoint(i, maxSize)
	}

	var population = make([]Agent, maxPopulation)
	for i := range population {
		population[i] = newRandomAgent(numPoints * 2)
	}

	// TODO: extract this into the agent.go file
	var pointCounts = make(map[int]int)
	for i := range points {
		pointCounts[i] = 0
	}

}
