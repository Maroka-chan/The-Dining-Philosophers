package main

import "sync"

type Fork struct {
	id        int
	timesUsed int
	inUse     int
	holder    *Philosopher
	input     chan int
	output    chan int
	arbiter   sync.Mutex
}

func (f *Fork) Run() {
	for {
		x := <-f.input
		if x == 1 {
			if f.inUse == 1 {
				f.output <- 0
			} else {
				f.inUse = 1
				f.timesUsed++
				f.output <- 1
			}
		} else if x == 2 {
			if f.inUse == 0 {
				f.output <- 0
			} else {
				f.inUse = 0
				f.output <- 1
			}
		} else if x == 10 {
			f.output <- f.timesUsed
		}
	}
}
