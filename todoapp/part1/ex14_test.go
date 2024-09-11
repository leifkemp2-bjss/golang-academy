package main

import (
	"fmt"
	"testing"
)

var ITERS = 100

func TestMain(m *testing.M){
	fmt.Printf("Performing benchmark with iteration size of %d\n", 2 * ITERS)
	m.Run()
}

func BenchmarkEx14(b *testing.B){
	for range b.N {
		ex14(ITERS)
	}
}

func BenchmarkEx15NoConcurrency(b *testing.B){
	for range b.N {
		ex15_noconcurrency(ITERS)
	}
}

func BenchmarkEx15Mutex(b *testing.B){
	for range b.N {
		ex15_mutex(ITERS)
	}
}

func BenchmarkEx15Channel(b *testing.B){
	for range b.N {
		ex15_channel(ITERS)
	}
}

func BenchmarkEx15ManyGoroutines(b *testing.B){
	for range b.N {
		ex15_manygoroutines(ITERS)
	}
}