package main

import "testing"

func BenchmarkPipeline(b *testing.B) {
	in, out := pipe(1000000)
	for i := 0; i < b.N; i++ {
		in <- 1
		<- out
	}
}
