package main

func search(nums []int, target int) int {
	// 原文提到升序排列，旋转之后分成两段, 但是两段内部还是升序排列
	// 二分查找
	left, right := 0, len(nums)-1
	for left <= right {
		mid := (right-left)/2 + left
		if nums[mid] == target {
			return mid
		}
		if nums[left] == target {
			return left
		}
		if nums[right] == target {
			return right
		}

		// 这里的等号
		if nums[left] < nums[mid] {
			// 这里的等号
			if nums[left] < target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			// 这里的等号
			if nums[mid] < target && target < nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}
	}
	return -1
}
