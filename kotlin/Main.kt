import kotlin.math.pow
import kotlin.math.sqrt
import kotlin.random.Random

// tuneable parameters
// TODO: extract these to somewhere nicer
const val NUM_POINTS = 256
const val MAX_SIZE = 10 // basically the dimensions of the world
const val MAX_POPULATION = 1000
const val NUM_GENERATIONS = 50
const val NUM_SURVIVORS = 100

typealias Point = Pair<Double, Double>

fun main() {
    val world =
            (0 until NUM_POINTS).map {
                Point(Random.nextDouble() * MAX_SIZE, Random.nextDouble() * MAX_SIZE)
            }

    var population = (0 until MAX_POPULATION).map { Agent() }.toMutableList()

    (0 until NUM_GENERATIONS).map {
        evalAndSort(population, world)
        val survivors = population.take(NUM_SURVIVORS)
        population = (0 until MAX_POPULATION).map { Agent(survivors) }.toMutableList()
    }

    evalAndSort(population, world)
}

fun distance(p1: Point, p2: Point): Double {
    return sqrt((p1.first - p2.first).pow(2.0) + (p1.second - p2.second).pow(2.0))
}

fun evalAndSort(population: MutableList<Agent>, world: List<Point>) {
    population.forEach { it.evaluate(world) }
    population.sortBy { it.score }
}
