package nsum

import "sort"

// 限制为正整数
func kSum2(nums []int, target int, k int) int {
	sort.Ints(nums)
	N := len(nums)
	dp := make([][][]int, N+1)
	for i := 0; i <= N; i++ {
		dp[i] = make([][]int, k+1)
		for j := 0; j <= k; j++ {
			dp[i][j] = make([]int, target+1)
		}
	}
	for i := 0; i <= N; i++ {
		dp[i][0][0] = 1
	}

	for i := 1; i <= N; i++ {
		for j := 1; j <= k && j <= i; j++ {
			for t := 1; t <= target; t++ {

				dp[i][j][t] = 0
				if t >= nums[i-1] {
					dp[i][j][t] = dp[i-1][j-1][t-nums[i-1]]
				}
				dp[i][j][t] += dp[i-1][j][t]

			}
		}

	}
	return dp[N][k][target]
}
