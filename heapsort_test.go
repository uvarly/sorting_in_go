package main

import (
	"testing"
)

func TestHeapsort(t *testing.T) {
	testSorting(t, heapsort)
}

func BenchmarkHeapsort(b *testing.B) {
	benchmarkSorting(b, heapsort)
}
