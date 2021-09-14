package main

import (
	"fmt"
	"sync"
	"time"
)

const PNUM int = 5

func main() {
	var forks [PNUM]*Fork
	var philosophers [PNUM]*Philosopher

	for i := 0; i < 5; i++ {
		forks[i] = &Fork{i, 0, 0, nil, make(chan int), make(chan int), sync.Mutex{}}
	}

	for i := 0; i < 5; i++ {
		philosophers[i] = &Philosopher{i, 0, forks[i], forks[(i+1)%PNUM], make(chan int), make(chan int)}
	}

	for _, fork := range forks {
		go fork.Run()
	}

	for i := 0; i < 4; i++ {
		go philosophers[i].QueryLoop()
		go philosophers[i].Run()
	}

	go philosophers[4].QueryLoop()
	go philosophers[4].Run_rev()

	// for _, philosopher := range philosophers {
	// 	go philosopher.QueryLoop()
	// 	go philosopher.Run()
	// }

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
