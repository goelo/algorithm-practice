package main

// O(1) 空间复杂度 摩尔投票法
func majorityElement(nums []int) int {
	var (
		major int
		count int
	)
	for _, num := range nums {
		if count == 0 {
			major = num
		}
		if major == num {
			count += 1
		} else {
			count -= 1
		}
	}
	return major
}
func majorityElement1(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	rec := make(map[int]int, len(nums))
	for i := 0; i < len(nums); i++ {
		rec[nums[i]]++
	}
	// 这里根据题目返回值，假设只有一个多数元素
	for _, v := range rec {
		if v > len(nums)/2 {
			return v
		}
	}
	return 0
}
