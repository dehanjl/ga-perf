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

	var genes = make([]int, geneSize)
	for i := range genes {
		genes[i] = rand.Intn(numPoints)
	}
	return Agent{genes, math.Inf(-1)}
}

// a new empty agent
func newAgent(geneSize int) Agent {
	return Agent{make([]int, geneSize), math.Inf(-1)}
}

func newAgentFromParents(parents []Agent) Agent {

	var newAgent = newAgent(len(parents[0].genes))
	// for each gene, choose a random parent
	for i := range newAgent.genes {
		newAgent.genes[i] = parents[rand.Intn(len(parents))].genes[i]
	}

	return newAgent
}

func (a *Agent) mutate() {
	// this will always mutate a gene
	// TODO: consider making this more interesting

	// choose a random gene
	var gene = rand.Intn(len(a.genes))

	// choose a random point
	var point = rand.Intn(numPoints)

	// set the gene to the point
	a.genes[gene] = point
}

// a function to evaluate the fitness of an agent
func (a *Agent) evaluate(points []Point) float64 {

	var score float64

	// mask to determine if all points have been visited
	var mask uint = determineMask(uint(numPoints)) // TODO: perhaps make this a parameter

	// keep track if all points have been visited
	// a uint is used basically used as a bit array
	// TODO: consider just using visitedCount
	var pointsCheck uint = 0

	var visitedCount = make([]int, numPoints)

	// set visitedCount and pointsCheck for the first gene
	visitedCount[a.genes[0]]++
	pointsCheck = setBit(pointsCheck, uint(a.genes[0]))

	i := 1 // start at 1 because we already "traversed" the first gene

	// traverse the path and calculate the distance
	for i < len(a.genes) {
		// set visitedCount and pointsCheck for the current gene
		visitedCount[a.genes[i]]++
		pointsCheck = setBit(pointsCheck, uint(a.genes[i])) // TODO: consider wrapping in an if

		score += distance(points[a.genes[i-1]], points[a.genes[i]])

		i++

		// if all points have been visited, break
		if checkMask(pointsCheck, mask) {
			break
		}
	}

	// add the distance back to the start
	score += distance(points[a.genes[i-1]], points[a.genes[0]])

	// for each point that has not been visited, add the distance to the closest point and back
	// TODO: This seems slow, consider just adding an arbitrary penalty
	for i := range visitedCount {
		if visitedCount[i] == 0 {
			var minDist float64 = math.Inf(1)
			for j := range visitedCount {
				if visitedCount[j] != 0 {
					var dist float64 = distance(points[i], points[j])
					if dist < minDist {
						minDist = dist
					}
				}
			}
			score += 2 * minDist
		}
	}

	a.score = score

	return a.score
}
