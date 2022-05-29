/*
 * @lc app=leetcode.cn id=1022 lang=golang
 *
 * [1022] 从根到叶的二进制数之和
 */

// @lc code=start
/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumRootToLeaf(root *TreeNode) (res int) {
	// 标准的 dfs
	var dfs func(node *TreeNode, curr int)
	dfs = func(node *TreeNode, curr int) {
		if node.Left != nil {
			dfs(node.Left, curr*2+node.Val)
		}
		if node.Right != nil {
			dfs(node.Right, curr*2+node.Val)
		}
		if node.Left == nil && node.Right == nil {
			res += curr*2 + node.Val
		}
	}
	if root.Left == nil && root.Right == nil {
		res = root.Val
		return
	}
	if root.Left != nil {
		dfs(root.Left, root.Val)
	}
	if root.Right != nil {
		dfs(root.Right, root.Val)
	}
	return
}

// @lc code=end

