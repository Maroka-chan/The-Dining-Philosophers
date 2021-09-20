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
}

func (p *Philosopher) Init(id int, left *Fork, right *Fork) {
	p.id = id
	p.left = left
	p.right = right
	p.input, p.output = make(chan int), make(chan int)
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
			p.left.input <- 1
			if <-p.left.output == 0 {
				continue
			}
			//fmt.Printf("P%d Query right fork %d\n", p.id, p.right.id)
			p.right.input <- 1
			if <-p.right.output == 0 {
				continue
			}

			fmt.Printf("Philosopher %d picks up Fork %d and %d\n", p.id, p.left.id, p.right.id)
			fmt.Printf("Philosopher %d is EATING\n", p.id)
			p.timesEaten++
			time.Sleep(time.Second * 5)

			p.left.input <- 2
			fmt.Println(<-p.left.input)

			p.right.input <- 2
			fmt.Println(<-p.right.input)

			fmt.Printf("Philosopher %d puts down Fork %d and %d\n", p.id, p.left.id, p.right.id)
			break
		}
	}
}

func (p *Philosopher) Run_rev() {
	for {
		fmt.Printf("Philosopher %d is THINKING\n", p.id)
		time.Sleep(time.Second)
		fmt.Printf("Philosopher %d is HUNGRY\n", p.id)
		for {
			//fmt.Printf("P%d Query right fork %d\n", p.id, p.right.id)
			p.right.input <- 1
			if <-p.right.output == 0 {
				continue
			}
			//fmt.Printf("P%d Query left fork %d\n", p.id, p.left.id)
			p.left.input <- 1
			if <-p.left.output == 0 {
				continue
			}

			fmt.Printf("Philosopher %d picks up Fork %d and %d\n", p.id, p.left.id, p.right.id)
			fmt.Printf("Philosopher %d is EATING\n", p.id)
			p.timesEaten++
			time.Sleep(time.Second)

			p.right.input <- 2
			fmt.Println(<-p.right.output)

			p.left.input <- 2
			fmt.Println(<-p.left.output)

			fmt.Printf("Philosopher %d puts down Fork %d and %d\n", p.id, p.left.id, p.right.id)
			break
		}
	}
}
