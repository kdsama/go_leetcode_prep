package grind75

import "math"

// 121. Best Time to Buy and Sell Stock

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func MaxProfit(prices []int) int {

	// {7,1,5,4,2}

	// lets try out greedy first
	sum := 0
	min := math.MaxInt32

	for i := range prices {
		if min > prices[i] {
			min = prices[i]
		}
		if sum < prices[i]-min {
			sum = prices[i] - min
		}
	}
	return sum
	// So greedy is causing

}
