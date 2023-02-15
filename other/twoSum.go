package main

func twoSum(nums []int, target int) []int {
	record := make(map[int]int)
	for i, v := range nums {
		// 没有记录过，记录下来
		if _, ok := record[v]; !ok {
			record[v] = i
		}
		// 如果存在两数之和，且不是同一个数
		if value, ok := record[target-v]; ok {
			if value != i {
				// 题目说了只有一个答案
				return []int{value, i}
			}
		}
	}
	return nil
}
