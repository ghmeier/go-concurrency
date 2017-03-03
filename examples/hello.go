package main

import (
	"fmt"
	"time"
)

func main() {
	// START OMIT
	go func() {
		fmt.Println("Hello World")
	}()
	// END OMIT
	time.Sleep(time.Duration(1) * time.Second)
	return
}
