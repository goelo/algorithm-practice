package sort

func quicksort(arr []int, left, right int) {
	if right <= left {
		return
	}
	j := partition(arr, left, right)
	quicksort(arr, left, j-1)
	quicksort(arr, j+1, right)
}

func partition(arr []int, left, right int) int {
	// 确定基准值
	pivot := arr[right]
	// 记录移动指针
	p := left
	for i := left; i < right; i++ {
		// 如果当前值比基准值小，和n[p]值交换，然后p移动
		if arr[i] < pivot {
			arr[i], arr[p] = arr[p], arr[i]
			p++
		}
	}
	// 最后在将n[p]和pivot值交换，这样就将比基准值小的数放在了左边，大的放在了右边
	arr[p], arr[right] = arr[right], arr[p]
	return p
}
