*Yes I am aware this "documentation" is sub-par, it is a work-in-progress*

# Genetic Algorithm Performance Benchmark

The idea is to implement various versions of a [Genetic Algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm) intended to solve the [Travelling Salesman Problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem). The goal is to implement the same algorithm in various langauages. in order to compare their performance. Furthermore this serves as an excuse to try multiple languages and learn how to use them as efficiently as possible.

# The World
The world consists of some number of points, and an agent must find the shortest path for visiting all points in this world.

# The Agents
Each agent contains an array of the order in which the points in the world are visited. For now, an agent is allowed to visit a point multiple times, and is punished for every point not visited. Furthermore, an agent does not need to return it's starting position. This heuristic algorithm is subject to change.

# The Implementations
## In Progess
- Golang
## Planned
- Python
- Rust
- Kotlin
