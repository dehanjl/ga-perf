use rand::{seq::SliceRandom, Rng};

// tuneable parameters
// TODO: extract these to somewhere nicer
const NUM_POINTS: i32 = 5;
const MAX_SIZE: i32 = 10; // basically the dimensions of the world
const MAX_POPULATION: i32 = 1000;
// const NUM_GENERATIONS: i32 = 50;
// const NUM_SURVIVORS: i32 = 100;

type Point = (f64, f64);

fn main() {
    let world: Vec<Point> = (0..NUM_POINTS).map(|_| new_point()).collect();
    let mut agents: Vec<Vec<i32>> = (0..MAX_POPULATION).map(|_| new_random_agent()).collect();

    agents.sort_by(|a, b| eval(a, &world).partial_cmp(&eval(b, &world)).unwrap());
}

fn new_point() -> Point {
    let x: f64 = rand::thread_rng().gen_range(0.0..(MAX_SIZE as f64));
    let y: f64 = rand::thread_rng().gen_range(0.0..(MAX_SIZE as f64));
    (x, y)
}

fn new_random_agent() -> Vec<i32> {
    let mut agent: Vec<i32> = (0..NUM_POINTS).collect();
    agent.shuffle(&mut rand::thread_rng());
    agent
}

fn new_agent(parent: &Vec<i32>) -> Vec<i32> {
    todo!("implement me")
}

fn distance(a: &Point, b: &Point) -> f64 {
    let dx = a.0 - b.0;
    let dy = a.1 - b.1;
    (dx * dx + dy * dy).sqrt()
}

fn eval(agent: &Vec<i32>, world: &Vec<Point>) -> f64 {
    let mut score = 0.0;
    for i in 0..NUM_POINTS {
        let j = (i + 1) % NUM_POINTS;
        score += distance(
            &world[agent[i as usize] as usize],
            &world[agent[j as usize] as usize],
        );
    }
    score
}
