/*
 * @lc app=leetcode.cn id=1325 lang=golang
 *
 * [1325] 删除给定值的叶子节点
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
func removeLeafNodes(root *TreeNode, target int) *TreeNode {
	if root == nil {
		return nil
	}
	root.Left = removeLeafNodes(root.Left, target)
	root.Right = removeLeafNodes(root.Right, target)

	if root.Left == nil && root.Right == nil && root.Val == target {
		return nil
	}
	return root
}

// @lc code=end

