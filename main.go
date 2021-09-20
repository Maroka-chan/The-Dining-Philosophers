package main

import (
	"fmt"
	"time"
)

const philosopherCount int = 5

const (
	NotEating = 0
	NotInUse  = 0
)

func main() {
	var forks [philosopherCount]*Fork
	for i := 0; i < philosopherCount; i++ {
		forks[i] = &Fork{i, 0, NotInUse, -1, make(chan int), make(chan int)}
	}

	var philosophers [philosopherCount]*Philosopher
	for i := 0; i < philosopherCount; i++ {
		leftIn, leftOut := make(chan int), make(chan int)
		rightIn, rightOut := make(chan int), make(chan int)

		philosophers[i] = &Philosopher{i, 0, NotEating, forks[i], forks[(i+1)%philosopherCount], make(chan int), make(chan int), leftIn, leftOut, rightIn, rightOut}
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
		for i := 0; i < 5; i++ {
			philosophers[i].input <- 1
			fmt.Printf("| P%d: %d ", i, <-philosophers[i].output)
		}
		fmt.Println("|")
		for i := 0; i < 5; i++ {
			forks[i].input <- 10
			fmt.Printf("| F%d: %d ", i, <-forks[i].output)
		}
		fmt.Println("|")
	}
}
