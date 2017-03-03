package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	c := make(chan string)
	// END1 OMIT

	//START2 OMIT
	go func() {
		fmt.Println(<-c)
	}()
	// END2 OMIT

	// START3 OMIT
	c <- "Hello World"
	// END3 OMIT
}
