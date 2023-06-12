package main

func minCostClimbingStairs(cost []int) int {
	// 楼顶不是在cost内部的，因此需要+1表示楼顶
	dp := make([]int, len(cost)+1)
	dp[0] = 0
	dp[1] = 0
	for i := 2; i < len(cost)+1; i++ {
		dp[i] = min1(dp[i-1]+cost[i-1], dp[i-2]+cost[i-2])
	}
	return dp[len(cost)]
}

func min1(a, b int) int {
	if a < b {
		return a
	}
	return b
}
