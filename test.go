package main

import (
	"fmt"
	"math/rand"
	"sort"
	"sync"
	"testing"
	"time"
)

var CASE_LEN = []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1_000, 10_000, 100_000}

var (
	testCases      []struct{ nums, sorted []int }
	benchmarkCases []struct{ nums []int }
)

func init() {
	var (
		mu sync.Mutex
		wg sync.WaitGroup
	)

	rand.Seed(time.Now().UnixNano())

	for _, l := range CASE_LEN {
		wg.Add(1)
		go func(l int) {
			var (
				sorted = make([]int, l)
				nums   = func() []int {
					nums := make([]int, 0, l)
					for i := 0; i < l; i++ {
						nums = append(nums, rand.Intn(l)-l/2)
					}
					return nums
				}()
			)

			defer wg.Done()
			copy(sorted, nums)
			sort.Slice(sorted, func(i, j int) bool { return sorted[i] < sorted[j] })

			mu.Lock()
			testCases = append(testCases, struct{ nums, sorted []int }{nums, sorted})
			benchmarkCases = append(benchmarkCases, struct{ nums []int }{nums})
			mu.Unlock()
		}(l)
	}

	wg.Wait()

	sort.Slice(testCases, func(i, j int) bool {
		return len(testCases[i].nums) < len(testCases[j].nums)
	})
	sort.Slice(benchmarkCases, func(i, j int) bool {
		return len(benchmarkCases[i].nums) < len(benchmarkCases[j].nums)
	})
}

func testSorting(t *testing.T, f func([]int) []int) {
	for _, tc := range testCases {
		var (
			nums   = make([]int, len(tc.nums))
			sorted = tc.sorted
		)

		copy(nums, tc.nums)

		t.Run(fmt.Sprintf("%d", len(nums)), func(t *testing.T) {
			var (
				expected = sorted
				received = f(nums)
			)

			t.Parallel()

			if len(expected) != len(received) {
				t.Errorf("array length modified; expected: %d, received: %d", len(expected), len(received))
				return
			}

			for i := range expected {
				if expected[i] != received[i] {
					t.Error("not sorted")
					return
				}
			}
		})
	}
}

func benchmarkSorting(b *testing.B, f func([]int) []int) {
	for _, bc := range benchmarkCases {
		nums := make([]int, len(bc.nums))

		copy(nums, bc.nums)

		b.Run(fmt.Sprintf("%d", len(nums)), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				f(nums)
			}
		})
	}
}
