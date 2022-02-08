package main

func bubble(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	endIdx := len(nums) - 1
	for endIdx > 0 {
		newEndIdx := 0
		for i := 0; i < endIdx; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				newEndIdx = i
			}
		}
		endIdx = newEndIdx
	}

	return nums
}
