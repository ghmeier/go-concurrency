package main

import (
	"fmt"
	// /"time"
)

type Counter struct {
	count int
	done  chan int
}

func main() {

	done := make(chan int)
	c := &Counter{count: 0, done: done}
	n := 10000
	for i := 0; i < n; i++ {
		go incrementN(c, n)()
	}

	finished := 0
	for finished < n {
		finished += <-done
	}

	fmt.Printf("result: %d\nexpect: %d\n", c.count, n*n)
	return
}

func incrementN(c *Counter, n int) func() {
	return func() {
		for i := 0; i < n; i++ {
			c.Count()
		}
		c.done <- 1
	}
}

func (c *Counter) Count() {
	c.count += 1
}
