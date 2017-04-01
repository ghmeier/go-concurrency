package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Printf("#%d\n", i)
}

func main() {
	done := make(chan int)
	// START OMIT
	for i := 0; i < 1000; i++ {
		go func(i int) {
			hello(i)  // HL
			done <- 1 // OMIT
		}(i)
	} // HL
	// END OMIT
	time.Sleep(time.Duration(1) * time.Second)
	total := 1
	for total < 1000 {
		total += <-done
	}
}
