package main

// 动态规划
// 重点是找到状态转移方程
func longestPalindrome(in string) string {
	if len(in) < 2 {
		return in
	}
	// dp[i][j] = s[i+1][j-1]
	dp := make([][]bool, len(in))
	for i := 0; i < len(in); i++ {
		dp[i] = make([]bool, len(in))
		dp[i][i] = true // 单个字符肯定是回文串
	}

	// 判断子串s[i:j] 是否是回文串即可，因此j天然>i,所以只需要考虑二维数组的对角线以上即可
	var (
		maxlen = 1
		start  int
	)
	for j := 1; j < len(in); j++ {
		// 找二维数组的右上半边
		// i含义代表前指针，所以i从0开始
		for i := 0; i < j; i++ {
			if in[i] != in[j] {
				dp[i][j] = false
			} else {
				// 主要判断边界条件, j-1-(i+1) +1 < 2 => j-i<3 表示去掉头尾之后的字符串剩余字符串无法构成回文字符串
				if j-i < 3 {
					dp[i][j] = true
				} else {
					// 	去掉头和尾，依然是回文字符串
					dp[i][j] = dp[i+1][j-1] //????
				}
			}
			// 只要 dp[i][j] == true 成立，就表示子串 s[i..j] 是回文，此时记录回文长度和起始位置
			if dp[i][j] && j-i+1 > maxlen {
				// abcmmcba
				maxlen = j - i + 1
				start = i
			}
		}
	}
	return in[start : start+maxlen]
}

func findLongestPalindrome(s string) string {
	// 中心扩散法
	return ""
}
