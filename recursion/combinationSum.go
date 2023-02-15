package recursion

// 深度遍历回溯
func combinationSum(candidates []int, target int) (ans [][]int) {
	var dfs func(target, idx int)
	combo := make([]int, 0)
	dfs = func(target, idx int) {
		if target == 0 {
			// 这里combo已经累加进结果里面了，所以下面的递归退出的时候要回滚。避免重复累加
			// 注意！！！这里不能直接用combo, 因为下面的combo改变，会影响这里的combo添加进ans的值。
			// ans = append(ans, combo)的写法是错误的。
			ans = append(ans, append([]int(nil), combo...))
			return
		}
		if idx == len(candidates) {
			return
		}
		// 不选i
		dfs(target, idx+1)
		// 选i
		if target-candidates[idx] >= 0 {
			combo = append(combo, candidates[idx])
			dfs(target-candidates[idx], idx)
			// 回溯看其他的分支
			combo = combo[:len(combo)-1]
		}
	}
	dfs(target, 0)
	return ans
}
