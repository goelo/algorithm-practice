package main

// 二分查找是常规解法
// 还有牛顿迭代法和袖珍计算器算法
func mySqrt(x int) int {
	// 只取整数部分，那么k*k <= x
	l, r := 0, x
	// 保存结果
	ans := -1
	for l <= r {
		// 避免溢出
		mid := l + (r-l)/2
		if mid*mid <= x {
			ans = mid
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return ans
}
