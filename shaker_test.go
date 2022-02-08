package main

import (
	"testing"
)

func TestShaker(t *testing.T) {
	testSorting(t, shaker)
}

func BenchmarkShaker(b *testing.B) {
	benchmarkSorting(b, shaker)
}
