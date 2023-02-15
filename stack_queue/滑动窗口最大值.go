package main

// 单调栈经典例题
type myQueue struct {
	que []int
}

func NewMyQueue() *myQueue {
	return &myQueue{
		que: make([]int, 0),
	}
}

func (q *myQueue) Empty() bool {
	return len(q.que) == 0
}

// 获取出口元素
func (q *myQueue) Front() int {
	return q.que[0]
}

// 获取入口元素
func (q *myQueue) Back() int {
	return q.que[len(q.que)-1]
}

func (q *myQueue) Push(value int) {
	// 当要push的元素，比队列入口的元素大的时候
	// 不停的把入口元素弹出, 当比入口元素小了，再push
	for !q.Empty() && value > q.Back() {
		q.que = q.que[:len(q.que)-1]
	}
	// 此时即将push的元素，比队列入口元素小
	q.que = append(q.que, value)
}

func (q *myQueue) Pop(value int) {
	if !q.Empty() && value == q.Front() {
		q.que = q.que[1:]
	}
}
func maxSlidingWindow(nums []int, k int) []int {
	queue := NewMyQueue()
	// 记录最大值
	res := make([]int, 0)
	// 填满窗口
	for i := 0; i < k; i++ {
		queue.Push(nums[i])
	}

	// ！！！此时要记录k滑动窗口的最大值
	res = append(res, queue.Front())
	// 接下来看窗口移动过程中数据的变化
	for i := k; i < len(nums); i++ {
		// 首先先把窗口最左侧，也就是出口的数值去掉
		// i位置是滑动窗口右侧即将移动到的位置，滑动窗口最左侧位置则为i-k，
		// 这个位置的值要pop出去
		queue.Pop(nums[i-k])
		// 把i位置的值push进来
		queue.Push(nums[i])
		// 找到当前滑动窗口的最大值，并记录下来
		res = append(res, queue.Front())
	}
	return res
}

func main() {

}
