package easy

func maxProfit(prices []int) int {
	if len(prices) == 0 || len(prices) == 1 {
		return 0
	}
	low, profit := prices[0], 0
	for _, v := range prices {
		if v < low {
			low = v
			continue
		}

		if v-low > profit {
			profit = v - low
		}
	}
	return profit
}
