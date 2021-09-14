package main

import "sync"

type Fork struct {
	id         int
	times_used int
	in_use     int
	holder     *Philosopher
	input      chan int
	output     chan int
	arbiter    sync.Mutex
}

func (f *Fork) Run() {
	for {
		x := <-f.input
		if x == 1 {
			if f.in_use == 1 {
				f.output <- 0
			} else {
				f.in_use = 1
				f.times_used++
				f.output <- 1
			}
		} else if x == 2 {
			if f.in_use == 0 {
				f.output <- 0
			} else {
				f.in_use = 0
				f.output <- 1
			}
		} else if x == 10 {
			f.output <- f.times_used
		}
	}
}
