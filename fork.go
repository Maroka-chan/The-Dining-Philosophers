package main

// Query enum
const (
	InUse = iota
	TimesUsed
)

const (
	PickUp  = iota + 2
	PutDown = iota + 2
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

		if x == PickUp {
			id := <-input
			if f.inUse == True {
				output <- False
			} else {
				f.inUse = True
				f.holderId = id
				output <- True
			}
		} else if x == PutDown {
			id := <-input
			forkWasUsed := <-input
			if f.inUse == False || id != f.holderId {
				output <- False
			} else {
				if forkWasUsed == True {
					f.timesUsed++
				}
				f.inUse = False
				output <- True
			}
		} else if x == TimesUsed {
			output <- f.timesUsed
		} else if x == InUse {
			output <- f.inUse
		}
	}
}
