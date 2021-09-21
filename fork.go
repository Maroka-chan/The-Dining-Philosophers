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

func (f *Fork) Run() {
	for {
		x := <-f.input

		if x == PickUp {
			id := <-f.input
			if f.inUse == True {
				f.output <- False
			} else {
				f.inUse = True
				f.holderId = id
				f.output <- True
			}
		} else if x == PutDown {
			id := <-f.input
			forkWasUsed := <-f.input
			if f.inUse == False || id != f.holderId {
				f.output <- False
			} else {
				if forkWasUsed == True {
					f.timesUsed++
				}
				f.inUse = False
				f.output <- True
			}
		} else if x == TimesUsed {
			f.output <- f.timesUsed
		} else if x == InUse {
			f.output <- f.inUse
		}
	}
}
