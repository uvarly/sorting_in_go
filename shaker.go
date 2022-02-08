package main

func shaker(nums []int) []int {
	beginIdx := 0
	endIdx := len(nums) - 1

	if len(nums) < 2 {
		return nums
	}

	for beginIdx <= endIdx {
		newBeginIdx := endIdx
		newEndIdx := beginIdx

		for i := beginIdx; i < endIdx; i++ {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				newEndIdx = i
			}
		}
		endIdx = newEndIdx

		for i := endIdx; i >= beginIdx; i-- {
			if nums[i] > nums[i+1] {
				nums[i], nums[i+1] = nums[i+1], nums[i]
				newBeginIdx = i
			}
		}
		beginIdx = newBeginIdx + 1
	}

	return nums
}
