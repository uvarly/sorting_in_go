package main

import (
	"testing"
)

func TestBubble(t *testing.T) {
	testSorting(t, bubble)
}

func BenchmarkBubble(b *testing.B) {
	benchmarkSorting(b, bubble)
}
