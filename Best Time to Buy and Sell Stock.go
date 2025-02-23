package main

func maxProfit(prices []int) int {
	profit := 0
	minProfit := prices[0]

	for i := 1; i < len(prices); i++ {
		if prices[i] < minProfit {
			minProfit = prices[i]
		} else if (prices[i] - minProfit) > profit {
			profit = prices[i] - minProfit
		}
	}
	return profit
}
