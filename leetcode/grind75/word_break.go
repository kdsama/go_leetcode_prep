package grind75

func wordBreak(s string, wordDict []string) bool {
	// problem currently with our code is branching is intense
	// can we skip a few branches ?
	// we can add a test case for letters as well ? if we there is a letter from s thats not present in wordDict, we break
	// we need to optimize the number of recursions that are happening.

	// it has to be done using dynamic programming
	mapping := map[string]bool{}
	for _, w := range wordDict {
		mapping[w] = true
	}

	dp := make([]bool, len(s)+1)
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for j := 0; j < i; j++ {
			if dp[j] && mapping[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}
