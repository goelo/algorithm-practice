package main

func generateParenthesis(n int) []string {
	var res = []string{}
	var dfs func(left, right, n int, path string)
	dfs = func(left, right, n int, path string) {
		// 左右括号都用光了，则退出递归
		if left == n && right == n {
			res = append(res, path)
		}
		// 左括号还未用光，left和right记录的是使用的个数。
		if left <= n {
			dfs(left+1, right, n, path+"(")
		}
		// 右括号未用光，并且左括号使用个数必须大于右括号，才可以加右括号
		if left > right && right <= n {
			dfs(left, right+1, n, path+")")
		}
	}
	dfs(0,0, n, "")
	return res
}

