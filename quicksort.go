package main

func partition(nums []int, lo, hi int) int {
	var (
		pivot = nums[(hi+lo)/2]
		i     = lo - 1
		j     = hi + 1
	)

	for {
		for i++; nums[i] < pivot; i++ {
		}

		for j--; nums[j] > pivot; j-- {
		}

		if i >= j {
			return j
		}

		nums[i], nums[j] = nums[j], nums[i]
	}
}

func _quicksort(nums []int, lo, hi int) []int {
	if lo < hi {
		pivot := partition(nums, lo, hi)

		_quicksort(nums, lo, pivot)
		_quicksort(nums, pivot+1, hi)
	}

	return nums
}

func quicksort(nums []int) []int {
	return _quicksort(nums, 0, len(nums)-1)
}
