package main

func lengthOfLongestSubstring(s string) (ans int) {
	start := 0
	m := make(map[byte]int)
	// 定义双指针，滑动窗口
	for end := 0; end < len(s); end++ {
		ch := s[end]
		// 如果存在, 移动start指针至最新的位置
		if index, ok := m[ch]; ok {
			// 这里记录start的最大的位置，例如abba, 这种情况下
			// 运行第二个a的时候，如果不做比较，则会取存储的index:0， start则从2变成了index+1=1
			start = max(start, index+1)
		}

		// 判断最大值
		ans = max(ans, end-start+1)
		// 更新记录的m
		m[ch] = end
	}
	return ans
}

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
