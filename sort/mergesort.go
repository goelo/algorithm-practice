package sort

import "fmt"

func mergeSort(arr []int) []int {
	// 递归求，首先判断递归退出条件
	if len(arr) == 1 {
		// 如果只剩下一个元素, 无需比较了，返回
		return arr
	}
	// 取中点
	mid := len(arr) / 2
	fmt.Println(mid)
	left := mergeSort(arr[:mid])
	right := mergeSort(arr[mid:])
	return merge(left, right)
}

func merge(left, right []int) []int {
	// 比较两个数组, 合并
	res := make([]int, 0)
	lIndex, rIndex := 0, 0
	// 有一个到达边界就需要退出
	for lIndex < len(left) && rIndex < len(right) {
		if left[lIndex] < right[rIndex] {
			res = append(res, left[lIndex])
			lIndex++
		} else {
			res = append(res, right[rIndex])
			rIndex++
		}
	}
	// 左边还未到达左数组边界
	if lIndex < len(left) {
		res = append(res, left[lIndex:]...)
	}
	if rIndex < len(right) {
		res = append(res, right[rIndex:]...)
	}
	return res
}
