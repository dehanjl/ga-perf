use noisy_float::prelude::r64;
use rand::{seq::SliceRandom, Rng};

// tuneable parameters
// TODO: extract these to somewhere nicer
const NUM_POINTS: i32 = 256;
const MAX_SIZE: i32 = 10; // basically the dimensions of the world
const MAX_POPULATION: i32 = 1000;
const NUM_GENERATIONS: i32 = 50;
const NUM_SURVIVORS: i32 = 100;

type Point = (f64, f64);
type World = Vec<Point>;
type Agent = Vec<i32>;
type Population = Vec<Agent>;

fn main() {
    // set up world
    let world: World = (0..NUM_POINTS).map(|_| new_point()).collect();

    // set up the initial population
    let mut population: Population = (0..MAX_POPULATION).map(|_| new_random_agent()).collect();

    for _ in 0..NUM_GENERATIONS {
        eval_and_sort(&mut population, &world);

        // determine the set of survivors
        let survivors = &population[0..NUM_SURVIVORS as usize];

        population = (0..MAX_POPULATION)
            .map(|_| new_agent_from_parents(survivors))
            .collect();
    }

    eval_and_sort(&mut population, &world);
}

fn new_point() -> Point {
    let x: f64 = rand::thread_rng().gen_range(0.0..MAX_SIZE as f64);
    let y: f64 = rand::thread_rng().gen_range(0.0..MAX_SIZE as f64);
    (x, y)
}

fn new_random_agent() -> Agent {
    let mut agent: Agent = (0..NUM_POINTS).collect();
    agent.shuffle(&mut rand::thread_rng());
    agent
}

fn new_agent_from_parents(parents: &[Agent]) -> Agent {
    let parent = parents.choose(&mut rand::thread_rng()).unwrap();

    let mut agent: Agent = parent.clone();
    agent.swap(
        rand::thread_rng().gen_range(0..NUM_POINTS) as usize,
        rand::thread_rng().gen_range(0..NUM_POINTS) as usize,
    );
    agent
}

fn distance(a: &Point, b: &Point) -> f64 {
    ((a.0 - b.0).powi(2) + (a.1 - b.1).powi(2)).sqrt()
}

fn eval(agent: &[i32], world: &[Point]) -> f64 {
    (0..NUM_POINTS)
        .zip((1..NUM_POINTS).chain(0..1))
        .map(|(i, j)| {
            distance(
                &world[agent[i as usize] as usize],
                &world[agent[j as usize] as usize],
            )
        })
        .sum()
}

fn eval_and_sort(pop: &mut [Agent], world: &[Point]) {
    pop.sort_by_cached_key(|a| r64(eval(a, world)));
}
