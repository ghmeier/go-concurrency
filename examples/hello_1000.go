package main

import (
	"fmt"
	"time"
)

func hello(i int) {
	fmt.Printf("#%d\n", i)
}

func main() {
	// START OMIT
	for i := 0; i < 1000; i++ { // HL
		go hello(i) // HL
	} // HL
	// END OMIT
	time.Sleep(time.Duration(1) * time.Second)
}
