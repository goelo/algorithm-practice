package sort

func HeapSort(arr []int) {
	if len(arr) < 2 {
		return
	}
	// 构建大顶堆 or 小顶堆
	for i := 0; i < len(arr); i++ {
		heapInsert(arr, i)
	}
	// 移除堆顶元素, 然后整理，保证堆顶始终满足大顶 or 小顶
	heapSize := len(arr)
	// 堆化
	for heapSize > 0 {
		arr[0], arr[heapSize-1] = arr[heapSize-1], arr[0]
		heapSize--
		heapify(arr, 0, heapSize)
	}
}

func heapInsert(arr []int, index int) {
	// 如果当前节点大于父节点的值
	for arr[index] > arr[(index-1)/2] {
		// 将当前节点移动到父节点
		arr[index], arr[(index-1)/2] = arr[(index-1)/2], arr[index]
		index = (index - 1) / 2
	}
}

// 主要是将当前index的数值来判断是否需要下沉，将比当前值大的数移动到顶端
func heapify(arr []int, index, heapSize int) {
	// 先看一下是否存在左孩子
	left := 2*index + 1
	// 左孩子有效前提下
	for left < heapSize {
		var larger int
		// 判断右孩子是否有效，然后判断左右孩子谁大，下标给larger
		if left+1 < heapSize && arr[left] < arr[left+1] {
			larger = left + 1
		} else {
			larger = left
		}
		//父和孩子，谁的值大，把下标给larger
		if arr[larger] < arr[index] {
			larger = index
		}
		if larger == index {
			break
		}
		// 交换
		arr[index], arr[larger] = arr[larger], arr[index]
		index = larger
		left = 2*index + 1
	}
}
