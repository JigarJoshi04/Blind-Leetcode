import "fmt"

func smallNumFinder(num1 int, num2 int) int {
	if num1 < num2 {
		return num1
	}

	return num2
}

func maxArea(height []int) int {
	lp := 0
	rp := len(height) - 1
	max := 0
	for lp < rp && lp >= 0 && rp <= len(height)-1 {
		w := (rp - lp) * smallNumFinder(height[lp], height[rp])
		if w > max {
			max = w
		}

		if height[lp] < height[rp] {
			lp = lp + 1
		} else {
			rp = rp - 1
		}
	}
	return max
}