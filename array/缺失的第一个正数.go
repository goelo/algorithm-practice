package main

import "fmt"

func firstMissingPositive(nums []int) int {
	for i := 0; i < len(nums); {
		idx := nums[i] - 1
		// nums[i]-1 作为idx必须有效，另外两者不能相等，此时交换位置
		if nums[i] > 0 && nums[i] != nums[idx] && nums[i] <= len(nums) {
			// i位置的值放在i位置值-1的位置上
			nums[i], nums[idx] = nums[idx], nums[i]
		} else {
			i++
		}
	}
	fmt.Println(nums)
	// 遍历完比对之后的nums，如果nums i位置上的值!= i+1，最先出现的即是最小的正数
	for idx, num := range nums {
		fmt.Println(idx)
		if num != idx+1 {
			return idx + 1
		}
	}
	// 都处理完都没有发现最小正数，则最小的正数为len(nums) + 1
	return len(nums) + 1
}
