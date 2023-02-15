package main

func minDistance(word1 string, word2 string) int {
	dp := make([][]int, len(word1)+1)
	for i := range dp {
		dp[i] = make([]int, len(word2)+1)
	}
	// word2 为空串
	for i := 1; i <= len(word1); i++ {
		dp[i][0] = i
	}
	// word1 为空串
	for j := 1; j <= len(word2); j++ {
		dp[0][j] = j
	}
	for i := 1; i <= len(word1); i++ {
		for j := 1; j <= len(word2); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1])
			} else {
				dp[i][j] = min(dp[i][j-1]+1, dp[i-1][j]+1, dp[i-1][j-1]+1)
			}
		}
	}
	return dp[len(word1)][len(word2)]
}

func min(nums ...int) int {
	min := nums[0]
	for _, v := range nums {
		if v < min {
			min = v
		}
	}
	return min
}

//func minDistance(word1 string, word2 string) int {
//
//	dp := [][]int{}
//
//	width := len(word1)
//	length := len(word2)
//
//	//初始化dp矩阵，多加了一个元素，对应dp[0][0]
//	for i := 0; i <= width; i++ {
//		dp = append(dp, make([]int, length+1))
//	}
//
//	for i := 1; i <= length; i++ {
//		dp[0][i] = i
//	}
//
//	for j := 1; j <= width; j++ {
//		dp[j][0] = j
//	}
//
//	for i := 1; i <= width; i++ {
//		for j := 1; j <= length; j++ {
//			if word1[i-1] == word2[j-1] {
//				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1])
//			} else {
//				dp[i][j] = min(dp[i-1][j]+1, dp[i][j-1]+1, dp[i-1][j-1]+1)
//			}
//		}
//	}
//
//	return dp[width][length]
//}
//
//func min(nums ...int) int {
//	min := nums[0]
//	for i := 1; i < len(nums); i++ {
//		if nums[i] < min {
//			min = nums[i]
//		}
//	}
//	return min
//}
