package main

import "sort"

func main() {
	arr := []int{-1, 0, 1, 2, -1, -4}
	threeSum(arr)
}

// 难点在于去重复的细节
// i的去重
func threeSum(nums []int) [][]int {
	// 可以直接使用排序方法
	sort.Ints(nums)
	res := make([][]int, 0)
	// 这里其实i的边界要把j,k的值去掉
	for i := 0; i < len(nums)-2; i++ {
		// 如果nums[i]都>0, 则无需判断了，直接退出。因为是升序排列
		if nums[i] > 0 {
			break
		}
		// 这里要判断后一个元素是否和前一个元素重复了，如果重复就跳过
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		j := i + 1
		k := len(nums) - 1
		for j < k {
			a, b, c := nums[i], nums[j], nums[k]
			if a+b+c == 0 {
				res = append(res, []int{a, b, c})
				// 这里要判断j和k的重复情况
				// 这里要持续判断，如果j和j的下一个元素重复的化，j要不听的右移跳过
				// !!! 这里的代码写法很优雅，但需要转个弯思考一下
				for j < k && b == nums[j] {
					j++
				}
				for j < k && c == nums[k] {
					k--
				}
			} else if a+b+c > 0 {
				k--
			} else {
				j++
			}
		}
	}
	return res
}
func quicksort1(nums []int, left, right int) {
	if left >= right {
		return
	}
	pivot := partition1(nums, left, right)
	quicksort1(nums, left, pivot-1)
	quicksort1(nums, pivot+1, right)
}

func partition1(nums []int, left, right int) int {
	// 先确定基准值
	pivot := nums[right]
	p := left
	for i := left; i < right; i++ {
		// 如果左边比基准值小，交换p和i的值
		if nums[i] < pivot {
			nums[i], nums[p] = nums[p], nums[i]
			p++
		}
	}
	// 确定第一个比基准值大的位置
	nums[right], nums[p] = nums[p], nums[right]
	return p
}
