package main

import (
	"fmt"
)

func hello(i int) {
	fmt.Printf("#%d\n", i)
}

func main() {
	done := make(chan bool)
	// START OMIT
	for i := 0; i < 1000; i++ {
		go func(i int) {
			hello(i)     // HL
			done <- true // OMIT
		}(i)
	} // HL
	// END OMIT
	total := 1
	for total < 1000 {
		if <-done {
			total++
		}
	}
	fmt.Println("Done")
}
