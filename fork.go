package main

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
				output <- 0
			} else {
				f.inUse = 1
				f.holderId = id
				f.timesUsed++
				output <- 1
			}
		} else if x == 2 {
			id := <-input
			if f.inUse == 0 || id != f.holderId {
				output <- 0
			} else {
				f.inUse = 0
				output <- 1
			}
		} else if x == 10 {
			output <- f.timesUsed
		}

	}
}
