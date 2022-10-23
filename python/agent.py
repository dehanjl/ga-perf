import random

from math import inf
from typing import List

from point import Point, distance


class Agent:
    def __init__(self, gene_size: int, parents: List['Agent'] = None):
        if parents is None:
            self.genes = list(range(gene_size))
            self.shuffle()
        else:
            self.crossover(parents[random.randint(0, len(parents) - 1)])
            self.mutate()

        self.score = inf

    def evaluate(self, points: List[Point]) -> float:
        # traverse the path, including back to the start
        self.score = sum(
            distance(points[i], points[j])
            for i, j in zip(self.genes, self.genes[1:] + self.genes[:1])
        )
        return self.score

    def shuffle(self):
        random.shuffle(self.genes)

    def mutate(self):
        i, j = random.sample(range(len(self.genes)), 2)
        self.genes[i], self.genes[j] = self.genes[j], self.genes[i]

    def crossover(self, parent: 'Agent'):
        self.genes = parent.genes[:]
