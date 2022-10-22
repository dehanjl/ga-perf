from typing import Tuple
from math import hypot
from random import uniform

Point = Tuple[float, float]


def distance(p1: Point, p2: Point) -> float:
    # pylint: disable=invalid-name
    return hypot(p2[0] - p1[0], p2[1] - p1[1])


def random_point(lim: float) -> Point:
    return uniform(0, lim), uniform(0, lim)
