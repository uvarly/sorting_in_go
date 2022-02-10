package main

import (
	"testing"
)

func TestCounting(t *testing.T) {
	testSorting(t, counting)
}

func BenchmarkCounting(b *testing.B) {
	benchmarkSorting(b, counting)
}
