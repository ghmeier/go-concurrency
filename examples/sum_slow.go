package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	sum := 0
	total := 100

	for i := 0; i < total; i++ {
		sum += retrieve(100)
	}
	fmt.Println(sum)
}

func retrieve(max int) int {
	time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)
	return rand.Intn(max)
}
