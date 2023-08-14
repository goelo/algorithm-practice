package main

// 使用双指针解法,主要是有序递增的数组，所以平方之后的最大值肯定是两边，不是中间
func sortedSquares(nums []int) []int {
	i, j, k := 0, len(nums)-1, len(nums)-1
	ans := make([]int, len(nums))

	for i <= j {
		a := nums[i] * nums[i]
		b := nums[j] * nums[j]
		// 两者的平方比较
		if a < b {
			ans[k] = b
			j--
		} else {
			ans[k] = a
			i++
		}
		k--
	}
	return ans
}
