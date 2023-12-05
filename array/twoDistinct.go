func containsDuplicate(nums []int) bool {
	mapping := make(map[int]int)
	dis := false
	for _, ele := range nums {
		_, ok := mapping[ele]
		if !ok {
			mapping[ele] = 1
		} else {
			dis = true
			return true
		}
	}
	return dis
}