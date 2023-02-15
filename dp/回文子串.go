package main

// https://www.bilibili.com/video/BV13Q4y197Wg/?vd_source=987a53a570f6253f01d0c41f4d4debdc
// https://leetcode.cn/problems/palindromic-substrings/solutions/917242/dai-ma-sui-xiang-lu-dai-ni-xue-tou-dpzi-vidge/
func countSubstrings(s string) int {
	var res int // 记录结果
	var countPalindromic func(start, end int)
	countPalindromic = func(start, end int) {
		// start和end不越界
		for start >= 0 && end < len(s) && s[start] == s[end] {
			res++
			// 往两边扩散
			start--
			end++
		}
	}
	// 中心扩散法
	// 遍历传入的字符串的每一位，确定每一个作为中心时候的结果
	for i := 0; i < len(s); i++ {
		// 当s有奇数个字符的时候中心点是一个
		countPalindromic(i, i)
		// 当s有偶数个字符的时候中心点有2个
		countPalindromic(i, i+1)
	}
	return res
}
