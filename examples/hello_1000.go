package main

import (
	"fmt"
)

func hello(i int) {
	if i == 999 {
		fmt.Printf("#%d\n", i)
	}
}

func main() {
	done := make(chan int, 1000)
	// START OMIT
	for i := 0; i < 1000; i++ {
		go func(i int) {
			hello(i)  // HL
			done <- 1 // OMIT
		}(i)
	} // HL
	// END OMIT
	total := 0
	for total < 1000 {
		total += <-done
	}
}
