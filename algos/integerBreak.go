package algos

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}

func IntegerBreak(n int) []int {
	dp := make([]int, n+1)
	dp[1] = 1

	for sum := 2; sum <= n; sum++ {
		if sum != n {
			dp[sum] = sum
		}
		for i := 1; i < n/2; i++ {
			if sum-i >= 0 {
				dp[sum] = max(dp[sum], dp[sum-i]*i)
			}
		}
	}

	return dp
}
