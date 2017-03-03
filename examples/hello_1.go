package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("Hello World")
}

func main() {
	// START OMIT
	go hello()
	// END OMIT
	time.Sleep(time.Duration(1) * time.Second)
	return
}
