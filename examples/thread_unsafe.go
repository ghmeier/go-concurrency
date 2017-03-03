package main

import (
	"fmt"
)

func main() {

	n, w := 100, 10000
	done := make(chan int)
	c := Counter(done)
	for i := 0; i < n; i++ {
		go c.increment(w)()
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
	done  chan int
}

func Counter(done chan int) *counter {
	return &counter{count: 0, done: done}
}

func (c *counter) increment(n int) func() {
	return func() {
		for i := 0; i < n; i++ {
			c.Count()
		}
		c.done <- 1
	}
}

func (c *counter) Count() {
	c.count += 1
}
