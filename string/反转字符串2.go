package main

func reverseStr(s string, k int) string {
	b := []byte(s)
	// 每隔2k个，那就是循环每次加2k个位置即可
	for i := 0; i < len(s); i += 2 * k {
		// 接下来只需要考虑每次的2k的前k个字符串反转即可
		// 但是需要注意的是，如果i+k超过了整个字符串的长度，那么就不能是i+k了，应该是len
		// 因为题目说了，如果剩余少于k个，则全部反转
		sub := b[i:min(i+k, len(b))]
		start := 0
		end := len(sub) - 1
		for start < end {
			sub[start], sub[end] = sub[end], sub[start]
			start++
			end--
		}
	}
	return string(b)
}
func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}
