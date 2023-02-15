package main

// 二分查找变种，主要思考如何确定边界
func searchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}
	left := findLeftPos(nums, target)
	// 如果左边未找到, 则直接返回，不需要找后面的右边界
	if left == -1 {
		return []int{-1, -1}
	}
	right := findRightPos(nums, target)
	return []int{left, right}
}

func findLeftPos(nums []int, target int) int {
	l := 0
	r := len(nums) - 1
	for l < r {
		// 确定下界
		mid := l + (r-l)/2
		// !!!这里注意 = target 和 > target的时候右边界的范围确认
		if nums[mid] >= target {
			// 说明mid右边的部分肯定不是第一次出现，此时搜索区间缩小为[left, mid]
			r = mid
		} else if nums[mid] < target {
			// 说明左边部分没有找到，left = mid +1
			l = mid + 1
		}
	}
	// for循环跳出条件是l >= r 其实是l=r即退出了，l不可能大于r
	// 所以需要判断当l=r时，是否找到target
	if nums[r] == target {
		return l
	}
	return -1
}

func findRightPos(nums []int, target int) int {
	l := 0
	r := len(nums)
	for l < r {
		// 确定上界
		mid := l + (r-l+1)/2
		if nums[mid] > target {
			// 大于target, 按照有序排列可以推出,target 在左半部分，此时[left, mid)的部分
			r = mid - 1
		} else if nums[mid] < target {
			// 小于target, 右半部分
			l = mid
		} else {
			// 相等的时候，需要重点判断
			// nums[mid]左边的元素一定不是target的最后一个元素, 因为nums[mid]也=target, 所以此时应该查找的区间是右边部分[mid, r]
			// 之所以不是mid+1，因为nums[mid]可能是最后一个元素, 所以不能去掉mid这个节点
			l = mid
		}
	}
	return r
}
