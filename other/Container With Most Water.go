package main

// 最优解 双指针 双指针可以用于确定边界问题
// 双指针代表的是 可以作为容器边界的所有位置的范围。
// 在一开始，双指针指向数组的左右边界，表示 数组中所有的位置都可以作为容器的边界
func maxArea(height []int) int {
	start, end := 0, len(height)-1
	var area int
	for start < end {
		area = max(area, min(height[start], height[end])*(end-start))
		if height[start] < height[end] {
			start++
		} else {
			end--
		}
	}
	return area
}

// 这样写法，时间复杂度超时
func maxArea1(height []int) int {
	var area int
	for i := 0; i < len(height); i++ {
		for j := 0; j < i; j++ {
			area = max(area, (i-j)*min(height[j], height[i]))
		}
	}
	return area
}
