import time
from typing import List
from agent import Agent
from point import random_point

# tuneable parameters
# TODO: extract these to somewhere nicer
NUM_POINTS = 256  # setting this over 63 might cause problems
MAX_SIZE = 10  # basically the dimensions of the world
MAX_POPULATION = 1000
NUM_GENERATIONS = 50
NUM_SURVIVORS = 100


def evaluate_population(pop: List['Agent']):
    for agent in pop:
        agent.evaluate(points)

    pop.sort(key=lambda x: x.score)

    return pop


if __name__ == '__main__':

    # define the points
    points = [random_point(MAX_SIZE) for _ in range(NUM_POINTS)]

    start_time = time.perf_counter_ns()

    # setup the first generation
    population = [Agent(NUM_POINTS) for _ in range(MAX_POPULATION)]

    # run the generations
    for _ in range(NUM_GENERATIONS):
        population = evaluate_population(population)

        # get the survivors
        survivors = population[:NUM_SURVIVORS]

        # make the next generation
        population = [Agent(NUM_POINTS, survivors)
                      for _ in range(MAX_POPULATION)]

    # evalutate the final population
    population = evaluate_population(population)

    end_time = time.perf_counter_ns()

    print("Best score:", population[0].score)

    print("Execution time:", (end_time - start_time) / 1000000, "ms")
    print("Time per generation:", (end_time - start_time) /
          1000000 / NUM_GENERATIONS, "ms")
