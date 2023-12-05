func maxSubArray(nums []int) int {
	max_sum := -100000
	new_arr := make([]int, len(nums))

	for index, ele := range nums {
		if index == 0 {
			new_arr[index] = ele
			max_sum = ele
		} else {
			prob_sum := new_arr[index-1] + ele
			if prob_sum > ele {
				new_arr[index] = prob_sum
			} else {
				new_arr[index] = ele
			}

			if new_arr[index] > max_sum {
				max_sum = new_arr[index]
			}

		}
	}

	return max_sum
}