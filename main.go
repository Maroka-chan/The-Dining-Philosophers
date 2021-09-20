package main

import (
	"fmt"
	"time"
)

const (
	False = iota
	True
)

const philosopherCount int = 5

func main() {
	var forks [philosopherCount]*Fork
	for i := 0; i < philosopherCount; i++ {
		forks[i] = &Fork{i, 0, False, -1, make(chan int), make(chan int)}
	}

	var philosophers [philosopherCount]*Philosopher
	for i := 0; i < philosopherCount; i++ {
		philosophers[i] = &Philosopher{i, 0, False, forks[i], forks[(i+1)%philosopherCount], make(chan int), make(chan int)}
	}

	for _, fork := range forks {
		go fork.Run()
	}

	for i := 0; i < philosopherCount; i++ {
		go philosophers[i].QueryLoop()
		go philosophers[i].Run()
	}

	for {
		time.Sleep(time.Second)
		for i := 0; i < philosopherCount; i++ {
			philosophers[i].input <- TimesEaten
			fmt.Printf("Philosopher %d:\n  times eaten: %d\n", i, <-philosophers[i].output)
			philosophers[i].input <- IsEating
			fmt.Printf("  eating: %d\n", <-philosophers[i].output)
			forks[i].input <- TimesUsed
			fmt.Printf("Fork %d:\n  times used: %d\n", i, <-forks[i].output)
			forks[i].input <- InUse
			fmt.Printf("  in use: %d\n", <-forks[i].output)
		}
	}
}
