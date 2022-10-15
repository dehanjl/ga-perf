*Yes I am aware this "documentation" is sub-par, it is a work-in-progress*

# Genetic Algorithm Performance Benchmark

The idea is to implement various versions of a [Genetic Algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm) intended to solve the [Travelling Salesman Problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem). The goal is to implement the same algorithm in various langauages. in order to compare their performance. Furthermore this serves as an excuse to try multiple languages and learn how to use them as efficiently as possible.

# The World
The world consists of some number of fixed points, and an agent must find the shortest path for visiting all points in this world.

# The Agents
## Type A Agents
These are the simple agents, where each agent contains an array of the order in which the points in the world are visited. For now, an agent is allowed to visit a point in the world multiple times.

The motivation for this type of agent is that it is the obvious solution to the problem.

## Type B Agents
These agents are sligthtly more complex. These agents contain a list of floating points which are possibly independent of the fixed points in the world. As the agents train, the floating points may move around, and the number may shift. If adding a fixed point in the world to the list would make the path more efficient, it is added.

The motivation for this type of agent is that it requires a more complex data structure and I hope that it will create a more satisfying visualization.

# The Heuristic Algorithm
The algorithms which is used to compute an Agent's score. 

## Type A Agents

0. Given: 
    - Track which if a point has been visited at all
    - Track the number of times each point is visited
1. Traverse the path calculate the distance for each step
    - If all points in the world has been visited, stop going through the path
2. Add the distance of going back to the first point
3. For each point which has not been visited
    - Calculate the distance to the nearest point that has been visited
    - Add the distance of going to that point and back


## Type B Agents

`TODO`

# The Implementations
For each language, it is intended to complete a first implementation without any fancy libraries or multi-processing. More advanced implementations making use of these techniques are also intended, in order to see how much performance can be gained.
## In Progess
- Golang (Type A)
## Planned
- Python
- Rust
- Kotlin

# Visualization
What is the point of creating an AI if you cant 