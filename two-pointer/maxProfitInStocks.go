func maxProfit(prices []int) int {
	lp := 0
	rp := 1
	max_profit := 0

	for rp < len(prices) {
		if (prices[rp] - prices[lp]) <= 0 {
			lp = rp
			rp += 1
		} else {
			if (prices[rp] - prices[lp]) > max_profit {
				max_profit = prices[rp] - prices[lp]
			}

			rp += 1
		}
	}

	return max_profit
}