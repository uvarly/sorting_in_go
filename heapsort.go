package main

func heapify(nums []int, i int) {
	var (
		largestIdx = i
		leftIdx    = 2*i + 1
		rightIdx   = 2*i + 2
	)

	if leftIdx < len(nums) && nums[leftIdx] > nums[largestIdx] {
		largestIdx = leftIdx
	}

	if rightIdx < len(nums) && nums[rightIdx] > nums[largestIdx] {
		largestIdx = rightIdx
	}

	if largestIdx != i {
		nums[largestIdx], nums[i] = nums[i], nums[largestIdx]
		heapify(nums, largestIdx)
	}
}

func heapsort(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := len(nums)/2 - 1; i >= 0; i-- {
		heapify(nums, i)
	}

	for i := len(nums) - 1; i >= 0; i-- {
		nums[0], nums[i] = nums[i], nums[0]
		heapify(nums[:i], 0)
	}

	return nums
}
