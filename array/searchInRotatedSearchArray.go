func searchItem(nums []int, lp int, rp int, target int) int {
	if lp > rp {
		return -1
	}
	mid := (lp + rp) / 2

	if nums[mid] == target {
		return mid
	}

	if nums[mid] < target {
		return searchItem(nums, mid+1, rp, target)
	} else {
		return searchItem(nums, lp, mid-1, target)
	}

	return -1
}

func pivotFinder(nums []int, lp int, rp int, target int) int {
	mid := (lp + rp) / 2

	if nums[lp] > nums[mid] {
		return pivotFinder(nums, lp, mid, target)
	}

	if nums[rp] < nums[mid] {
		return pivotFinder(nums, mid+1, rp, target)
	}

	return lp
}

func search(nums []int, target int) int {
	pivotIndex := pivotFinder(nums, 0, len(nums)-1, target)
	itemIndex1 := searchItem(nums, 0, pivotIndex, target)
	indexItem2 := searchItem(nums, pivotIndex, len(nums)-1, target)

	if itemIndex1 == -1 {
		return indexItem2
	} else {
		return itemIndex1
	}
	return -1
}

