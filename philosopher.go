package main

import (
	"time"
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

func (p *Philosopher) Init(id int, left *Fork, right *Fork) {
	p.id = id
	p.left = left
	p.right = right
	p.input, p.output = make(chan int), make(chan int)
	p.leftIn, p.leftOut = make(chan int), make(chan int)
	go p.left.Run(p.leftIn, p.leftOut)
	p.rightIn, p.rightOut = make(chan int), make(chan int)
	go p.right.Run(p.rightIn, p.rightOut)
}

func (p *Philosopher) QueryLoop() {
	for {
		x := <-p.input
		switch x {
		case 0:
			p.output <- p.timesEaten
		case 1:
			p.output <- p.eating
		}
	}
}

func (p *Philosopher) Run() {
	for {
		time.Sleep(time.Second)
		for {
			p.leftIn <- 1
			p.leftIn <- p.id
			if <-p.leftOut == 0 {
				time.Sleep(time.Second * 5)
				continue
			}
			p.rightIn <- 1
			p.rightIn <- p.id
			if <-p.rightOut == 0 {
				p.leftIn <- 2
				p.leftIn <- p.id
				<-p.leftOut
				time.Sleep(time.Second * 5)
				continue
			}

			p.eating = 1
			time.Sleep(time.Second * 5)
			p.eating = 0
			p.timesEaten++

			p.leftIn <- 2
			p.leftIn <- p.id
			<-p.leftOut

			p.rightIn <- 2
			p.rightIn <- p.id
			<-p.rightOut
			time.Sleep(time.Second * 1)
			break
		}
	}
}
