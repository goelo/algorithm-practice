package main

func permute(nums []int) (ans [][]int) {
	var visit []int // 表示用过的数字
	var dfs func(tmp []int, k int)
	dfs = func(tmp []int, k int) {
		// 判断什么时候退出
		if k == len(nums) {
			ans = append(ans, append([]int(nil), tmp...))
			return
		}
		for i := 0; i < len(nums); i++ {
			// 判断这个数字是否填过。
			if isVisit(visit, nums[i]) {
				continue
			}
			// 标记这个数已经填过
			visit = append(visit, nums[i])
			tmp = append(tmp, nums[i])
			// 继续判断下一个位置
			dfs(tmp, k+1)
			// 然后回溯标记的数字
			visit = visit[:len(visit)-1]
			tmp = tmp[:len(tmp)-1]
		}
		return
	}

	dfs([]int{}, 0)
	return ans
}

func isVisit(nums []int, target int) bool {
	for _, num := range nums {
		if num == target {
			return true
		}
	}
	return false
}
