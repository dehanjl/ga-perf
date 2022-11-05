# Aknowledgements
[@keegancsmith](https://github.com/keegancsmith)
[@mvniekerk](https://github.com/mvniekerk)

# Genetic Algorithm Performance Benchmark

The idea is to implement various versions of a [Genetic Algorithm](https://en.wikipedia.org/wiki/Genetic_algorithm) intended to solve the [Travelling Salesman Problem](https://en.wikipedia.org/wiki/Travelling_salesman_problem). The goal is to implement the same algorithm in various langauages. in order to compare their performance. Furthermore this serves as an excuse to try multiple languages and learn how to use them as efficiently as possible.

# The World
The world consists of some number of fixed points, and an agent must find the shortest path for visiting all points in this world.

# The Agents
## Type 1 Agents
These are the simple agents, where each agent contains an array of the order in which the points in the world are visited. The initial generation of the agents is initialised with a shuffled array of all the points in the world. Thus an agent will always visit all the points in the world. An agent is mutated by swapping two genes.

The motivation for this type of agent is that it is the obvious solution to the problem.

## Type 2 Agents
These agents are sligthtly more complex. These agents contain a list of floating points which are possibly independent of the fixed points in the world. As the agents train, the floating points may move around, and the number may shift. If adding a fixed point in the world to the list would make the path more efficient, it is added.

The motivation for this type of agent is that it requires a more complex data structure and I hope that it will create a more satisfying visualization.

# The Heuristic Algorithm
The algorithms which is used to compute an Agent's score. 

## Type 1 Agents

0. Given: 
    - Track which if a point has been visited at all
    - Track the number of times each point is visited
1. Traverse the path calculate the distance for each step
    - If all points in the world has been visited, stop going through the path
2. Add the distance of going back to the first point
3. For each point which has not been visited
    - Calculate the distance to the nearest point that has been visited
    - Add the distance of going to that point and back


## Type 2 Agents

`TODO`

# The Implementations
For each language, it is intended to complete a first implementation without any fancy libraries or multi-processing. More advanced implementations making use of these techniques are also intended, in order to see how much performance can be gained.
## Basic Level of Completion
- Golang (Type 1)
- Python (Type 1)
- Rust (Type 1)
## In Progess
## Planned
- Kotlin
- Clojure

# Visualization
What is the point of creating an AI if you cant create a cool GIF for twitter? Currently I am looking at the following ways of animating the learning process:
1. [SFML](https://www.sfml-dev.org/) with C++
    - Probably a pain in the arse, but fun.
2. [Manim](https://github.com/ManimCommunity/manim/) with Python
    - I can pretend I'm Grant Sanderson.
3. *Something* with Typescript
    - I guess I need to practice this as well.

# Performance Comparison
In order to compare the different implementations, it is intended to measure all implementations relative to some baseline implementation.
## Benchmarking Settings
### Standard Settings
*This is subject to change*
- Number of points in the world = 256
- Number of agents per generation = 1000
- Number of generations = 50
- Number of survivors per generation = 100
- Number of runs = 16
## Results
# Results from using [Hyperfine](https://github.com/sharkdp/hyperfine)
- Rust: 1x
- Golang: 2.67x
- Kotlin (JVM): 13.71x
- Python: 49.71x
