package main

import (
	"math"
)

func counting(nums []int) []int {
	var (
		min = math.MaxInt
		max = math.MinInt
	)

	if len(nums) < 2 {
		return nums
	}

	for _, v := range nums {
		if v > max {
			max = v
		}

		if v < min {
			min = v
		}
	}

	var (
		offset = max - min
		counts = make([]int, offset+1)
		i      = 0
	)

	for _, v := range nums {
		counts[v-min]++
	}

	for j, v := range counts {
		for ; v > 0; v-- {
			nums[i] = j + min
			i++
		}
	}

	return nums
}
