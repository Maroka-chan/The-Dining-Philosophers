package main

import (
	"fmt"
)

type Fork struct {
	id        int
	timesUsed int
	inUse     int
	holderId  int
	input     chan int
	output    chan int
}

func (f *Fork) Run(input chan int, output chan int) {
	for {
		x := <-input

		if x == 1 {
			id := <-input
			if f.inUse == 1 {
				fmt.Printf("Philosopher %d failed to pick up fork %d because philosopher %d is holding it\n", id, f.id, f.holderId)
				output <- 0
			} else {
				fmt.Printf("Philosopher %d managed to pick up fork %d\n", id, f.id)
				f.inUse = 1
				f.holderId = id
				f.timesUsed++
				output <- 1
			}
		} else if x == 2 {
			id := <-input
			if f.inUse == 0 || id != f.holderId {
				fmt.Printf("Philosopher %d failed to put down fork %d because philosopher %d is holding it\n", id, f.id, f.holderId)
				output <- 0
			} else {
				fmt.Printf("Philosopher %d managed to put down fork %d because philosopher %d was holding it\n", id, f.id, f.holderId)
				f.inUse = 0
				output <- 1
			}
		} else if x == 10 {
			output <- f.timesUsed
		}

	}
}
