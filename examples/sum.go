package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {

	for i := 0; i < 2; i++ {
		select {
		case o := <-results(retrieve(100)):
			fmt.Printf("Finished 1: %d\n", o)
		case o := <-results(slowRetrieve(1000)):
			fmt.Printf("Finished 2: %d\n", o)
		}
	}
}

func results(f func() int) <-chan int {
	c, sum := make(chan int), make(chan int)
	go combine(c, sum, 10)
	for i := 0; i < 10; i++ {
		go func() { c <- f() }()
	}
	return sum
}

func combine(c, result chan int, max int) {
	sum := 0
	for i := 0; i < max; i++ {
		sum += <-c
	}
	result <- sum
}

func retrieve(max int) func() int {
	return func() int {
		time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
		return rand.Intn(max)
	}
}

func slowRetrieve(max int) func() int {
	return func() int {
		time.Sleep(time.Duration(rand.Intn(5000)) * time.Millisecond)
		return rand.Intn(max)
	}
}
