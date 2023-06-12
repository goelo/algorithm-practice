package main

import (
	"math/rand"
	"time"
)

// 时间复杂度为O(N), 一般用于top k问题。从海量数字中找到第k大的数
// 将数组分成大于k和小于k的两组，小于k的有n-k个, 大于k的有k-1个。第K大的数就是第n-k个数。
func FindLargestK(nums []int, k int) int {
	// 生成随机种子
	rand.Seed(time.Now().UnixNano())
	// 这里是len(nums)-k， 而不是传入K，因为最终要找的是第n-k个数，就是第K大的数
	return quickSelect(nums, 0, len(nums)-1, len(nums)-k)
}
func quickSelect(nums []int, l, r int, target int) int {
	// 分区
	i := randomPartition(nums, l, r)
	if i == target {
		return nums[i]
	} else if i < target {
		return quickSelect(nums, i+1, r, target)
	}
	return quickSelect(nums, l, i-1, target)
}

func randomPartition(arr []int, left, right int) int {
	index := rand.Int()%(right-left+1) + left
	// 确定基准值
	arr[index], arr[right] = arr[right], arr[index]
	// 记录移动指针
	i := left
	for k := left; k < right; k++ {
		// 如果当前值比基准值小，和n[p]值交换，然后p移动
		if arr[k] < arr[right] {
			arr[k], arr[i] = arr[i], arr[k]
			i++
		}
	}
	// 最后在将n[p]和pivot值交换，这样就将比基准值小的数放在了左边，大的放在了右边
	arr[i], arr[right] = arr[right], arr[i]
	return i
}
