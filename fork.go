package main

import "sync"

type Fork struct {
	arbiter    sync.Mutex
	times_used int
	in_use     int
	input      chan int
	output     chan int
}

func (f *Fork) init() {
	f.times_used = 0
	f.in_use = 0
	f.input, f.output = make(chan int), make(chan int)
}

func (f *Fork) Run() {
	for {
		x := <-f.input
		f.arbiter.Lock()
		if x == 1 {
			f.output <- f.in_use
		} else if x == 2 {
			f.output <- f.times_used
		}
		f.arbiter.Unlock()
	}
}

func (f *Fork) PickUp() {
	f.arbiter.Lock()
	if f.in_use == 0 {
		f.times_used++
		f.in_use = 1
	}
	f.arbiter.Unlock()
}

func (f *Fork) PutDown() {
	f.arbiter.Lock()
	f.in_use = 0
	f.arbiter.Unlock()
}
