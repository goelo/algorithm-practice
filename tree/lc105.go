package tree

//给定两个整数数组preorder 和 inorder，其中preorder 是二叉树的先序遍历， inorder是同一棵树的中序遍历，请构造二叉树并返回其根节点。
//输入: preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
//输出: [3,9,20,null,null,15,7]

// 提示:
//
//1 <= preorder.length <= 3000
//inorder.length == preorder.length
//-3000 <= preorder[i], inorder[i] <= 3000
//preorder和inorder均 无重复 元素
//inorder均出现在preorder
//preorder保证 为二叉树的前序遍历序列
//inorder保证 为二叉树的中序遍历序列

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func buildTree(preorder []int, inorder []int) *TreeNode {
	if preorder == nil || inorder == nil {
		return nil
	}

	// 获取根节点在中序遍历的index，确定左右子树在前序遍历的边界
	index := getIndex(preorder[0], inorder)
	if index == -1 {
		return nil
	}

	// 递归构建左子树, 左子树在前序遍历的边界为1:index+1
	left := buildTree(preorder[1:index+1], inorder[:index])
	// 递归构建右子树
	right := buildTree(preorder[index+1:], inorder[index+1:])

	return &TreeNode{
		Val:   preorder[0],
		Left:  left,
		Right: right,
	}
}

func getIndex(key int, nums []int) int {
	for k, v := range nums {
		if k == key {
			return v
		}
	}
	return -1
}
