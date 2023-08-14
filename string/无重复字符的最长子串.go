package main

// 双指针解法
func lengthOfLongestSubstring1(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	// "abcabcbb"
	num := 0
	left, right := 0, 0
	for right < len(s) {
		// 这里要把右边界前面所有的字符，都和右边界指向的字符比较，而不是仅仅比较左边界和右边界的字符
		for i := left; i < right; i++ {
			// 所以循环判断，left指针需要从左边界移动到右边界，每个字符和右边界的字符比较
			if s[i] == s[right] {
				// 记录长度
				num = max(num, right-left)
				// 左边界移动到当前相等字符串i位置的后一位
				left = i + 1
			}
		}
		right++
	}
	// 这里要考虑一种情况就是，内层的循环，始终没有处理
	// 例如“au”，因此要再次计算num的结果
	num = max(num, right-left)

	return num
}
func max1(a, b int) int {
	if a > b {
		return a
	}
	return b
}
