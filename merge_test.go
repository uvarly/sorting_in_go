package main

import (
	"testing"
)

func TestMerge(t *testing.T) {
	testSorting(t, merge)
}

func BenchmarkMerge(b *testing.B) {
	benchmarkSorting(b, merge)
}
