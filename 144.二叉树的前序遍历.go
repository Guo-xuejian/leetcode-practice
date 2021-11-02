/*
 * @lc app=leetcode.cn id=144 lang=golang
 *
 * [144] 二叉树的前序遍历
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
func preorderTraversal(root *TreeNode) (res []int) {
	// 二叉树的遍历在函数内部实现方法去递归
	var preOrder func(*TreeNode)
	preOrder = func(node *TreeNode) {
		if node != nil { // 只要节点不为空就写入
			res = append(res, node.Val)
			preOrder(node.Left)
			preOrder(node.Right)
		} else {
			return
		}
	}
	preOrder(root)
	return
}

// @lc code=end

