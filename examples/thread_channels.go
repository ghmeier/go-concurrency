package main

import (
	"fmt"
)

func main() {

	n, w := 100, 10000
	done := make(chan int)
	c := Counter(done)
	for i := 0; i < n; i++ {
		go c.incrementN(w)()
	}

	finished := 0
	for finished < n {
		finished += <-done
	}

	fmt.Printf("result: %d\nexpect: %d\n", c.count, n*w)
	return
}

type counter struct {
	count int
	in    chan int
	done  chan int
}

func Counter(done chan int) *counter {
	c := &counter{
		count: 0,
		done:  done,
		in:    make(chan int),
	}
	c.start()
	return c
}

func (c *counter) incrementN(n int) func() {
	return func() {
		for i := 0; i < n; i++ {
			c.Count()
		}

		c.done <- 1
	}
}

func (c *counter) start() {
	go func() {
		for {
			c.count += <-c.in
		}
	}()
}

func (c *counter) Count() {
	c.in <- 1
}
