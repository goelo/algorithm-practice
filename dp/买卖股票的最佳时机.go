package main

func maxProfit(prices []int) int {
	// 利润最大化 如果用动态规划解首先考虑状态转移方程如何设计
	// dp[i] 为第i天买卖股票的最大利润, 那么需要考虑两种情况：
	// 1. 第i天交易，获得最大利润, 交易则需要记录历史最低价是多少，dp[i] = prices[i] - minPrice(历史最低价)
	// 2. 第i天不交易，那么最大利润则是前一天的dp[i] = dp[i-1]
	// 综合dp[i] = max(prices[i]-minPrice, dp[i-1])
	dp := make([]int, 0, len(prices))
	// 需要注意i要大于1
	dp[0] = 0
	minPrice := prices[0]
	for i := 1; i < len(prices); i++ {
		// 找到历史最低价
		minPrice = min(minPrice, prices[i])
		dp[i] = max(dp[i-1], prices[i]-minPrice)
	}
	// 这里index=最后一个是因为一旦确定最大利润之后，则不会改变
	return dp[len(prices)-1]
}
