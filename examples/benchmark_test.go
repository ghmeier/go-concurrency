package main

import (
	"fmt"
	"testing"
)

func BenchmarkLogger(b *testing.B) {
	errLog := make(chan string)
	stdLog := make(chan string)
	go handle(b.N, errLog, stdLog)()

	for i := 0; i < b.N; i++ {
		errLog <- fmt.Sprintf("%d", i)
		stdLog <- fmt.Sprintf("%d", i)
	}
}

func handle(n int, err, std <-chan string) func() {
	return func() {
		for i := 0; i < n*2; i++ {
			select {
			case s := <-err:
				fmt.Printf("ERROR: %s\n", s)
			case s := <-std:
				fmt.Println(s)
			}
		}
	}
}
