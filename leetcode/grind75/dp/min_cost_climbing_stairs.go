package dp

func MinCostClimbingStairs(cost []int) int {

	m := len(cost)
	dp := make([]int, m+1)
	dp[0] = cost[0]
	dp[1] = cost[1]
	for i := 2; i < m; i++ {
		dp[i] = cost[i] + min(dp[i-1], dp[i-2])

	}
	dp[m] = min(dp[m-1], dp[m-2])
	return dp[m]
}
