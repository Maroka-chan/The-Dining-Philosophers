package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id         int
	timesEaten int
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
		if x == 1 {
			p.output <- p.timesEaten
		}
	}
}

func (p *Philosopher) Run() {
	for {
		fmt.Printf("Philosopher %d is THINKING\n", p.id)
		time.Sleep(time.Second)
		fmt.Printf("Philosopher %d is HUNGRY\n", p.id)
		for {
			//fmt.Printf("P%d Query left fork %d\n", p.id, p.left.id)
			fmt.Printf("Philosopher %d tries to pick up Fork %d and %d\n", p.id, p.left.id, p.right.id)
			p.leftIn <- 1
			p.leftIn <- p.id
			if <-p.leftOut == 0 {
				//fmt.Printf("Philosopher %d failed\n", p.id)
				time.Sleep(time.Second * 5)
				continue
			}
			//fmt.Printf("P%d Query right fork %d\n", p.id, p.right.id)
			p.rightIn <- 1
			p.rightIn <- p.id
			if <-p.rightOut == 0 {
				fmt.Printf("Philosopher %d failed and is putting down fork %d\n", p.id, p.left.id)
				p.leftIn <- 2
				p.leftIn <- p.id
				fmt.Println(<-p.leftOut)
				time.Sleep(time.Second * 5)
				continue
			}

			fmt.Printf("Philosopher %d succceded\n", p.id)
			fmt.Printf("Philosopher %d is EATING\n", p.id)
			p.timesEaten++
			time.Sleep(time.Second * 5)

			fmt.Printf("Philosopher %d puts down Fork %d and %d\n", p.id, p.left.id, p.right.id)

			p.leftIn <- 2
			p.leftIn <- p.id
			fmt.Println(<-p.leftOut)

			p.rightIn <- 2
			p.rightIn <- p.id
			fmt.Println(<-p.rightOut)
			time.Sleep(time.Second * 1)
			break
		}
	}
}
