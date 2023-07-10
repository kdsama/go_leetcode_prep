package grind75

import "math"

func coinChange(coins []int, amount int) int {

	if amount == 0 {
		return 0
	}
	minCoin := math.MaxInt
	for i := range coins {
		if minCoin >= coins[i] {
			minCoin = coins[i]
		}
	}

	if minCoin > amount {
		return -1
	}
	dp := make([]int, amount+1)
	dp[0] = 1

	for i := 1; i <= amount; i++ {
		mi := math.MaxInt
		for _, coin := range coins {
			if i-coin >= 0 && dp[i-coin] >= 0 {
				if mi > dp[coin]+dp[i-coin] {
					mi = dp[coin] + dp[i-coin]
				}
			}
		}
		if mi == math.MaxInt {
			dp[i] = -1
		} else {
			dp[i] = mi
		}
	}

	return dp[amount]

}
