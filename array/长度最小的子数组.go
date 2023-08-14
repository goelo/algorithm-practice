package main

// 滑动窗口
func minSubArrayLen(target int, nums []int) int {
	var (
		sum int
		j   int
	)
	ans := len(nums) + 1
	// 窗口右边界移动
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 判断是否移动窗口左边
		for sum >= target {
			sublength := i - j + 1
			if sublength < ans {
				ans = sublength
			}
			sum -= nums[j]
			j++
		}
	}
	if ans == len(nums)+1 {
		return 0
	}
	return ans
}
