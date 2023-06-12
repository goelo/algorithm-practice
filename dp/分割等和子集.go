package main

// 还原为0-1背包的问题，如果忘记了。看代码随想录解答
func canPartition(nums []int) bool {
	// 1.clarify what's dp meaning
	// dp[j]容量为j的背包，所背的物品价值最大可以为dp[j]
	// 2. 确定递推公式
	// dp[i] = max(dp[i], dp[i-nums[i]]+nums[i])确认
	// 3.初始化dp
	// 4. 确定遍历顺序
	// 此题中让找到是否存在可分割成两个数组，数组之和相等
	// 则转换为数组a+数组b的累计和 = nums的累计和
	// 5. 举例推导dp
	var (
		target int
		sum    int
	)
	for _, num := range nums {
		sum += num
	}

	// 不是偶数，则不能找到
	if sum%2 != 0 {
		return false
	}
	target = sum / 2
	// 这里的个数是根据题意动态分析的
	dp := make([]int, target+1)
	for i := 0; i < len(nums); i++ {
		// 这里的倒序要额外注意！！！
		for j := target; j >= nums[i]; j-- {
			dp[j] = max(dp[j], dp[j-nums[i]]+nums[i])
			// fmt.Println(dp[j])
		}
	}

	return dp[target] == target
}
