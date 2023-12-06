func minRecur(nums []int, lp int, rp int) int {
	mid := (lp + rp) / 2
	// 0, mid =1
	if lp == rp {
		return nums[lp]
	}
	if rp-lp == 1 {
		if nums[lp] < nums[rp] {
			return nums[lp]
		} else {
			return nums[rp]
		}
	}
	if nums[lp] > nums[mid] {
		return minRecur(nums, lp, mid)
	} else if nums[rp] < nums[mid] {
		return minRecur(nums, mid, rp)
	}

	return nums[0]
}

func findMin(nums []int) int {
	return minRecur(nums, 0, len(nums)-1)
}

