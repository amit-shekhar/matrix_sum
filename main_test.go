package main

import "testing"

func BenchmarkMatrix1(b *testing.B) {
	prepare()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		singleThread()
	}
}

func BenchmarkMatrix4(b *testing.B) {
	prepare()
	// run the Fib function b.N times
	for n := 0; n < b.N; n++ {
		fourThreads()
	}
}
