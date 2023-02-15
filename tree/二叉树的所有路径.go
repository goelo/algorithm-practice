package tree

import "strconv"

// 更优雅的解法
func treePaths(root *TreeNode) (res []string) {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		// 到达叶子节点
		if root.Left == nil && root.Right == nil {
			// 这个时候把节点追加进去
			v := path + strconv.Itoa(root.Val)
			res = append(res, v)
			return
		}
		// 不是叶子节点
		// 这里最后是->, 这样便于回溯, 只有到达叶子节点才会追加节点值进去
		path += strconv.Itoa(root.Val) + "->"
		if root.Left != nil {
			dfs(root.Left, path)
		}
		if root.Right != nil {
			dfs(root.Right, path)
		}
	}
	dfs(root, "")
	return res
}

// 自己解法,看过答案
func binaryTreePaths(root *TreeNode) (res []string) {
	if root == nil {
		return nil
	}
	var dfs func(root *TreeNode, path string)
	dfs = func(root *TreeNode, path string) {
		if root.Left == nil && root.Right == nil {
			res = append(res, path)
			return
		}
		if root.Left != nil {
			tmp := path
			tmp += "->" + strconv.Itoa(root.Left.Val)
			dfs(root.Left, tmp)
		}
		if root.Right != nil {
			tmp := path
			tmp += "->" + strconv.Itoa(root.Right.Val)
			dfs(root.Right, tmp)
		}
		return
	}
	dfs(root, strconv.Itoa(root.Val))
	return res
}
