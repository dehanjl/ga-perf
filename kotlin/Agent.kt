import kotlin.random.Random

class Agent {
    var genes: MutableList<Int>
    var score = Double.POSITIVE_INFINITY

    /** Create a new agent with random genes */
    constructor(geneSize: Int = NUM_POINTS) {
        genes = (0 until geneSize).toList().shuffled().toMutableList()
    }

    /** Create a new agent from a list of parents */
    constructor(parents: List<Agent>) {
        genes = parents[Random.nextInt(parents.size)].genes.toMutableList()
        mutate()
    }

    private fun mutate() {
        val i = Random.nextInt(genes.size)
        val j = Random.nextInt(genes.size)
        genes[i] = genes[j].also { genes[j] = genes[i] }
    }

    fun evaluate(world: List<Point>) {
        (world.indices)
                .zip(world.indices.drop(1) + 0)
                .sumOf { (i, j) -> distance(world[genes[i]], world[genes[j]]) }
                .let { score = it }
    }
}
