package main

import (
	"fmt"
	"time"
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
	time.Sleep(time.Duration(1) * time.Second)
}

func handle(err, std <-chan string) func() {
	return func() {
		// START4 OMIT
		for {
			select {
			case s := <-err: // HL
				fmt.Printf("ERROR: %s\n", s)
			case s := <-std: // HL
				fmt.Println(s)
			}
		}
		// END4 OMIT
	}
}
