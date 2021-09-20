package main

import (
	"fmt"
	"time"
)

type Philosopher struct {
	id          int
	times_eaten int
	left        *Fork
	right       *Fork
	input       chan int
	output      chan int
	left_in     chan int
	left_out    chan int
	right_in    chan int
	right_out   chan int
}

func (p *Philosopher) Init(id int, left *Fork, right *Fork) {
	p.id = id
	p.left = left
	p.right = right
	p.input, p.output = make(chan int), make(chan int)
	p.left_in, p.left_out = make(chan int), make(chan int)
	go p.left.Run(p.left_in, p.left_out)
	p.right_in, p.right_out = make(chan int), make(chan int)
	go p.right.Run(p.right_in, p.right_out)
}

func (p *Philosopher) QueryLoop() {
	for {
		x := <-p.input
		if x == 1 {
			p.output <- p.times_eaten
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
			p.left_in <- 1
			p.left_in <- p.id
			if <-p.left_out == 0 {
				//fmt.Printf("Philosopher %d failed\n", p.id)
				time.Sleep(time.Second * 5)
				continue
			}
			//fmt.Printf("P%d Query right fork %d\n", p.id, p.right.id)
			p.right_in <- 1
			p.right_in <- p.id
			if <-p.right_out == 0 {
				fmt.Printf("Philosopher %d failed and is putting down fork %d\n", p.id, p.left.id)
				p.left_in <- 2
				p.left_in <- p.id
				fmt.Println(<-p.left_out)
				time.Sleep(time.Second * 5)
				continue
			}

			fmt.Printf("Philosopher %d succceded\n", p.id)
			fmt.Printf("Philosopher %d is EATING\n", p.id)
			p.times_eaten++
			time.Sleep(time.Second * 5)

			fmt.Printf("Philosopher %d puts down Fork %d and %d\n", p.id, p.left.id, p.right.id)

			p.left_in <- 2
			p.left_in <- p.id
			fmt.Println(<-p.left_out)

			p.right_in <- 2
			p.right_in <- p.id
			fmt.Println(<-p.right_out)
			time.Sleep(time.Second * 1)
			break
		}
	}
}
