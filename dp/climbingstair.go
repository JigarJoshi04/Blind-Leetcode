
func climbStairs(n int) int {
	a := 1
	b := 2
	c := a + b

	if n == 1 {
		return a
	}

	if n == 2 {
		return b
	}

	if n == 3 {
		return c
	}
	for i := 0; i < n-3; i++ {
		a = b
		b = c
		c = a + b
	}

	return c
}