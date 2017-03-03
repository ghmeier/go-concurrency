package main

import (
	"fmt"
)

func main() {
	// START1 OMIT
	errLog := make(chan string)
	stdLog := make(chan string)
	// END1 OMIT

	//START2 OMIT
	go handle(errLog, stdLog)()
	// END2 OMIT

	// START3 OMIT
	errLog <- "Unable to Connect"
	stdLog <- "Successfully Connected"

	stdLog <- "Received new request"
	errLog <- "Request failed"
	// END3 OMIT
}

func handle(err, std <-chan string) func() {
	return func() {
		// START4 OMIT
		for {
			fmt.Printf("ERROR: %s\n", <-err) // HL
			fmt.Println(<-std)               // HL
		}
		// END4 OMIT
	}
}
