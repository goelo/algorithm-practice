package main

func maximalRectangle(matrix [][]byte) (ans int) {
	if len(matrix) == 0 {
		return 0
	}
	// 统计'1'的在ij位置上的值
	left := make([][]int, 0, len(matrix))
	for i, row := range matrix {
		// 固定长度
		left[i] = make([]int, 0, len(matrix[0]))
		for j, v := range row {
			if v == '0' {
				// 如果是'0', 无需统计，跳过
				continue
			}
			// 记录'1'的累积值
			if j == 0 {
				// 上面'0'判断拦截了，所以此处的j肯定为'1'
				left[i][j] = 1
			} else {
				left[i][j] = left[i][j-1] + 1
			}
		}
	}

	// 计算最小矩形的面积
	for i, row := range matrix {
		for j, v := range row {
			if v == '0' {
				continue
			}
			// 记录宽度
			width := left[i][j]
			// 面积， 高为1
			area := width
			// 从i往上查找, 找面积最大的点
			for k := i - 1; k >= 0; k-- {
				width = min(width, left[k][j])
				// 宽*长
				area = max(area, width*(i-k+1))
			}
			// 找最大的记录
			ans = max(ans, area)
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
