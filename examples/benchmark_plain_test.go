package main

import (
	"fmt"
	"testing"
)

func BenchmarkPlain(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fmt.Printf("ERROR: %d\n", i)
		fmt.Printf("%d\n", i)
	}
}
