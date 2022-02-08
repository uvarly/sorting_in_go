package main

import (
	"testing"
)

func TestInsertion(t *testing.T) {
	testSorting(t, insertion)
}

func BenchmarkInsertion(b *testing.B) {
	benchmarkSorting(b, insertion)
}
