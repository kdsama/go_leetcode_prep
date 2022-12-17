package grind75

// 121. Best Time to Buy and Sell Stock

// You are given an array prices where prices[i] is the price of a given stock on the ith day.

// You want to maximize your profit by choosing a single day to buy one stock and choosing a different day in the future to sell that stock.

// Return the maximum profit you can achieve from this transaction. If you cannot achieve any profit, return 0.

func MaxProfit(prices []int) int {

	// {7,1,5,4,2}
	// We need to find min value and max value
	// we need to make sure min value index is before max value
	// if there is no profit return 0
	// Any two needs to be compared apart from themselves. So n^2 might happen, but thats not what we want anyways
	// How to do it in n ?
	// while min_position < max_position
	// if next position yields a smaller item , thats our new min position
	// same goes for max position
	//  constraint ==> sum > 0
	return 1
}
