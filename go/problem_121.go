package main

func maxProfit(prices []int) int {
	profit := 0
	currentLowest := prices[0]
	for i := 1; i < len(prices); i++ {
		if prices[i] < currentLowest {
			currentLowest = prices[i]
		} else if prices[i]-currentLowest > profit {
			profit = prices[i] - currentLowest
		}
	}
	return profit
}
