package main

import (
	"testing"
)

func TestQuicksort(t *testing.T) {
	testSorting(t, quicksort)
}

func BenchmarkQuicksort(b *testing.B) {
	benchmarkSorting(b, quicksort)
}
