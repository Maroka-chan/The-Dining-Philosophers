package main

import (
	"fmt"
)

type Fork struct {
	id         int
	times_used int
	in_use     int
	holder_id  int
	input      chan int
	output     chan int
}

func (f *Fork) Run(input chan int, output chan int) {
	for {
		x := <-input

		if x == 1 {
			id := <-input
			if f.in_use == 1 {
				fmt.Printf("Philosopher %d failed to pick up fork %d because philosopher %d is holding it\n", id, f.id, f.holder_id)
				output <- 0
			} else {
				fmt.Printf("Philosopher %d managed to pick up fork %d\n", id, f.id)
				f.in_use = 1
				f.holder_id = id
				f.times_used++
				output <- 1
			}
		} else if x == 2 {
			id := <-input
			if f.in_use == 0 || id != f.holder_id {
				fmt.Printf("Philosopher %d failed to put down fork %d because philosopher %d is holding it\n", id, f.id, f.holder_id)
				output <- 0
			} else {
				fmt.Printf("Philosopher %d managed to put down fork %d because philosopher %d was holding it\n", id, f.id, f.holder_id)
				f.in_use = 0
				output <- 1
			}
		} else if x == 10 {
			//fmt.Printf("Fork %d was queried about times used\n", f.id)
			output <- f.times_used
		}

	}
}
