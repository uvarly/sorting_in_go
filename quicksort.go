package main

func partition(nums []int, lo, hi int) int {
	pivot := nums[(hi+lo)/2]
	i, j := lo-1, hi+1

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
