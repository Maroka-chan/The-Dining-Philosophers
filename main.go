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
		leftIn, leftOut := make(chan int), make(chan int)
		rightIn, rightOut := make(chan int), make(chan int)

		philosophers[i] = &Philosopher{i, 0, False, forks[i], forks[(i+1)%philosopherCount], make(chan int), make(chan int), leftIn, leftOut, rightIn, rightOut}
		go philosophers[i].left.Run(leftIn, leftOut)
		go philosophers[i].right.Run(rightIn, rightOut)
	}

	for _, fork := range forks {
		go fork.Run(fork.input, fork.output)
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
