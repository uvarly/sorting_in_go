package main

func selection(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	for i := 0; i < len(nums)-1; i++ {
		minIdx := i
		for j := i + 1; j < len(nums); j++ {
			if nums[j] < nums[minIdx] {
				minIdx = j
			}
		}
		nums[i], nums[minIdx] = nums[minIdx], nums[i]
	}

	return nums
}
