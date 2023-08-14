package main

func maxAreaOfIsland(grid [][]int) int {
	var area int
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == 1 {
				area = max(area, dfs(grid, i, j))
			}
		}
	}
	return area
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
func dfs(grid [][]int, i, j int) int {
	// 判断边界, 因为递归，所以当i=0时，递归进来是<0的。
	// 还要判断是0的岛屿
	if i < 0 || j < 0 || i >= len(grid) || j >= len(grid[0]) || grid[i][j] == 0 {
		return 0
	}
	// 记录走过的岛屿
	grid[i][j] = 0
	return 1 + dfs(grid, i-1, j) + dfs(grid, i+1, j) + dfs(grid, i, j-1) + dfs(grid, i, j+1)
}
