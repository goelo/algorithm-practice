package main

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// 动态规划
func longestValidParentheses(s string) int {
	// 假设dp[i]表示以s[i]为最后一位时有效子串的长度, 这里要注意有效两个字！！！，就是必须以i位置结尾可以有效。
	// 如果s[i] == (, 则根本找不到之前与之匹配的括号，因此是无效的。dp[i]因此为0, 这点如果不能想通，这道题就会卡住。
	maxAns := 0
	// 初始化dp
	dp := make([]int, len(s))
	// 注意这里的i是从1开始的。因为0位置dp[0] = 0
	for i := 1; i < len(s); i++ {
		// 只有)才有效，当s[i]为(时，dp[i]为0.因为按照dp[i]的定义，s[i]必须能匹配才符合有效，否则就要重新开始计算。
		if s[i] == ')' {
			// 判断i-1位置是) or (
			if s[i-1] == '(' {
				// 当i超过2位置时，i-2才有效
				if i >= 2 {
					dp[i] = dp[i-2] + 2
				} else {
					// 只剩下i-1和i位置组成的有效()
					dp[i] = 2
				}
				// 注意这里判断的是i对称位置是否有效，即>0
			} else if i-dp[i-1] > 0 && s[i-dp[i-1]-1] == '(' {
				// 如果i位置对称的位置是(, 则计算dp[i]
				// 这个时候同样需要判断i-dp[i-1]是否超过2的位置, 因为这样i的对称位置才有效
				if i-dp[i-1] >= 2 {
					dp[i] = dp[i-1] + 2 + dp[i-dp[i-1]-2]
				} else {
					// 如果不超过2位置，则dp
					dp[i] = dp[i-1] + 2
				}
			}
			maxAns = max(maxAns, dp[i])
		}
	}
	return maxAns
}
