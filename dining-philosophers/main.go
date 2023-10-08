package main

import (
	"fmt"
	"sync"
	"time"
)

// The Dining Philosophers problem is a classic synchronization and concurrency problem that illustrates challenges related to resource allocation and deadlock prevention. Here's a concise description of the problem in 7 lines:

// In a dining room, there are five philosophers sitting around a circular table.
// Each philosopher spends their time thinking and eating.
// To eat, a philosopher needs two forks, one on their left and one on their right.
// Philosophers can only pick up one fork at a time and must put both forks down after eating.
// To avoid deadlock, philosophers follow the "resource hierarchy" rule, which means they pick up the fork with the lower number first.
// However, this can lead to a situation where all philosophers pick up the fork on their left simultaneously, causing a deadlock.
// To solve this problem, various synchronization techniques like semaphores or mutexes are used to ensure that philosophers can safely pick up forks and eat without leading to deadlock.

// This is a simple implementation of Dijkistra solution to the "Dining Philosophers" dilemma

// Philosopher is a struct which stores an information about the philosopher
type Philosopher struct {
	name string
	rightFork int
	leftFork int
}

// philosopher list of all philosophers

var philosophers = []Philosopher {
	{name: "Plato", leftFork: 4, rightFork: 0},
	{name: "Socrates", leftFork: 0, rightFork: 1},
	{name: "Aristotle", leftFork: 1, rightFork: 2},
	{name: "Pascal", leftFork: 2, rightFork: 3},
	{name: "Locke", leftFork: 3, rightFork: 4},
}


var hunger = 3 // how many times does a person eat
var eatTime = 1 * time.Second
var thinkTime = 3 * time.Second
var sleepTime = 1 * time.Second

func main() {
	// print out a welcome message
	fmt.Println("Dining Philosophers Problem")
	fmt.Println("======================================")
	fmt.Println("The table is empty")

	// start the meal
	dine()


	// print out results
	fmt.Println("The table is empty")
}

type ForkMap = map[int] *sync.Mutex

func dine() {
	wg := &sync.WaitGroup{}
	wg.Add(len(philosophers))
	
	seated := &sync.WaitGroup{}
	seated.Add(len(philosophers))

	// forks is a map of all 5 forks
	forks := make(ForkMap)
	for i := 0; i < len(philosophers); i++ { 
		forks[i] = &sync.Mutex{}
	}
	
	// start the meal
	for i := 0; i < len(philosophers); i++ {
		// fire off a goroutine for the current philosopher
		go diningProblem(philosophers[i], wg, forks, seated)
	}

	wg.Wait()
}

func diningProblem(philosopher Philosopher, wg *sync.WaitGroup, forks ForkMap, seated *sync.WaitGroup) {
	defer wg.Done()

	// seat philosopher at the table
	fmt.Printf("%s is seated at the table./n", philosopher.name)
	seated.Done()

	// eat three times
	for i := hunger; i > 0; i --{
		// get a lock on both forks
		forks[philosopher.leftFork].Lock()
		fmt.Printf("%s takes the left fork.\n", philosopher.name)
		forks[philosopher.rightFork].Lock()
		fmt.Printf("%s takes the right fork.\n", philosopher.name)
	}
}