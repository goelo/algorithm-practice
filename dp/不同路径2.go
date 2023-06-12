package main

// 这题需要意识到如果是障碍物的话，dp[i][j]=0表示没有路径可以到达。
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	// 需要考虑边界条件
	// 如果第一步和最后一步是障碍，那么没有路径可达
	if obstacleGrid[0][0] == 1 || obstacleGrid[m-1][n-1] == 1 {
		return 0
	}
	// 初始化dp
	dp := make([][]int, m)
	for i, _ := range dp {
		dp[i] = make([]int, n)
	}

	// 初始化, 如果是障碍物, 后面的就都是0, 不用循环了
	// !!!注意代码里for循环的终止条件，一旦遇到obstacleGrid[i][0] == 1的情况就停止dp[i][0]的赋值1的操作，dp[0][j]同理
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 有障碍物，跳过。跳过即dp[i][j] = 0
			if obstacleGrid[i][j] == 1 {
				continue
			}
			dp[i][j] = dp[i][j-1] + dp[i-1][j]
		}
	}
	return dp[m-1][n-1]
}
