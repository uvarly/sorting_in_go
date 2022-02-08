package main

import (
	"testing"
)

func TestSelection(t *testing.T) {
	testSorting(t, selection)
}

func BenchmarkSelection(b *testing.B) {
	benchmarkSorting(b, selection)
}
