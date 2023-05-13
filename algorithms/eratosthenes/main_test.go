package main

import "testing"

// benchmark

func BenchmarkSimple(b *testing.B) {
	simple(b.N)
}

func BenchmarkEratosthenes(b *testing.B) {
	eratosthenes(b.N)
}
