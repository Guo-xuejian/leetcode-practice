/*
 * @lc app=leetcode.cn id=145 lang=golang
 *
 * [145] 二叉树的后序遍历
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
func postorderTraversal(root *TreeNode) (res []int) {
	var postOrder func(*TreeNode)
	postOrder = func(node *TreeNode) {
		if node != nil { // 后序遍历顺序，左子树、右子树、根节点
			postOrder(node.Left)
			postOrder(node.Right)
			res = append(res, node.Val)
		} else {
			return
		}
	}

	postOrder(root)
	return
}

// @lc code=end

