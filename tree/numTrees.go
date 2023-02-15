package tree

func numTrees(n int) int {
	// 0-n个
	dp := make([]int, n+1)
	dp[0] = 1
	// dp[i] 表示i个节点组成不同bst的组合个数
	// 求dp[n] 可以理解为枚举树根节点x x从1-n
	// 那么dp[x] = 左子树组合数dp[x-1]*右子树组合数dp[n-x]

	// 枚举节点数目
	for i := 1; i <= n; i++ {
		// 计算i个节点能构造的bst的数量dp[i]
		for j := 1; j <= i; j++ {
			// 枚举根节点个数
			dp[i] += dp[j-1] * dp[i-j]
		}
	}
	return dp[n]
}
