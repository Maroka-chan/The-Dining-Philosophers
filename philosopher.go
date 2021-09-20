package main

import (
	"time"
)

// Query enum
const (
	TimesEaten = iota
	IsEating
)

type Philosopher struct {
	id         int
	timesEaten int
	eating     int
	left       *Fork
	right      *Fork
	input      chan int
	output     chan int
	leftIn     chan int
	leftOut    chan int
	rightIn    chan int
	rightOut   chan int
}

func (p *Philosopher) QueryLoop() {
	for {
		x := <-p.input
		switch x {
		case TimesEaten:
			p.output <- p.timesEaten
		case IsEating:
			p.output <- p.eating
		}
	}
}

func (p *Philosopher) Run() {
	for {
		time.Sleep(time.Second)
		for {
			p.leftIn <- PickUp
			p.leftIn <- p.id
			if <-p.leftOut == False {
				time.Sleep(time.Second * 5)
				continue
			}
			p.rightIn <- PickUp
			p.rightIn <- p.id
			if <-p.rightOut == False {
				p.leftIn <- PutDown
				p.leftIn <- p.id
				p.leftIn <- False
				<-p.leftOut
				time.Sleep(time.Second * 5)
				continue
			}

			p.eating = True
			time.Sleep(time.Second * 5)
			p.eating = False
			p.timesEaten++

			p.leftIn <- PutDown
			p.leftIn <- p.id
			p.leftIn <- True
			<-p.leftOut

			p.rightIn <- PutDown
			p.rightIn <- p.id
			p.leftIn <- True
			<-p.rightOut
			time.Sleep(time.Second * 1)
			break
		}
	}
}
