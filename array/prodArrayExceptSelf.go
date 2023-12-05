func productExceptSelf(nums []int) []int {
	product := 1

	ct := 0
	for _, ele := range nums {
		if ele != 0 {
			product *= ele
		} else {
			ct += 1
		}
	}

	ans := make([]int, len(nums))
	for index, ele := range nums {
		if ct == 0 {
			ans[index] = product / (ele)
		} else {
			if ct > 1 {
				ans[index] = 0
			} else if ele != 0 {
				ans[index] = 0
			} else {
				ans[index] = product
			}
		}
	}

	return ans
}