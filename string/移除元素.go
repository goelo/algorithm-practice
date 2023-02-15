package main

func removeElement(nums []int, val int) int {
	var i int
	for i < len(nums) {
		if nums[i] == val {
			nums = append(nums[0:i], nums[i+1:]...)
			i = 0
			continue
		}
		i++

	}
	return len(nums)
}
