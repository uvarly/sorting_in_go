package main

func join(nums []int, lo, mi, hi int) {
	var (
		nums1   = make([]int, mi-lo+1)
		nums2   = make([]int, hi-mi)
		i, j, k = 0, 0, lo
	)

	copy(nums1, nums[lo:mi+1])
	copy(nums2, nums[mi+1:hi+1])

	for ; i < len(nums1) && j < len(nums2); k++ {
		if nums1[i] <= nums2[j] {
			nums[k] = nums1[i]
			i++
		} else {
			nums[k] = nums2[j]
			j++
		}
	}

	for ; i < len(nums1); i, k = i+1, k+1 {
		nums[k] = nums1[i]
	}

	for ; j < len(nums2); j, k = j+1, k+1 {
		nums[k] = nums2[j]
	}
}

func _merge(nums []int, lo, hi int) []int {
	if lo >= hi {
		return nums
	}

	mi := (lo + hi) / 2

	_merge(nums, lo, mi)
	_merge(nums, mi+1, hi)

	join(nums, lo, mi, hi)

	return nums
}

func merge(nums []int) []int {
	if len(nums) < 2 {
		return nums
	}

	return _merge(nums, 0, len(nums)-1)
}
