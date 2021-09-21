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
			p.left.input <- PickUp
			p.left.input <- p.id
			if <-p.left.output == False {
				time.Sleep(time.Second * 5)
				continue
			}
			p.right.input <- PickUp
			p.right.input <- p.id
			if <-p.right.output == False {
				p.left.input <- PutDown
				p.left.input <- p.id
				p.left.input <- False
				<-p.left.output
				time.Sleep(time.Second * 5)
				continue
			}

			p.eating = True
			time.Sleep(time.Second * 5)
			p.eating = False
			p.timesEaten++

			p.left.input <- PutDown
			p.left.input <- p.id
			p.left.input <- True
			<-p.left.output

			p.right.input <- PutDown
			p.right.input <- p.id
			p.right.input <- True
			<-p.right.output
			time.Sleep(time.Second * 1)
			break
		}
	}
}
