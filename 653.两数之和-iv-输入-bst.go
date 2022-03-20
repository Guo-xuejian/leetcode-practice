/*
 * @lc app=leetcode.cn id=653 lang=golang
 *
 * [653] 两数之和 IV - 输入 BST
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
func findTarget(root *TreeNode, k int) bool {
	numMap := map[int]struct{}{}
	var dfs func(*TreeNode) bool
	dfs = func(node *TreeNode) bool {
		if node == nil {
			return false
		}
		if _, ok := numMap[node.Val]; ok {
			return true
		}
		numMap[k - node.Val] = struct{}{}
		return dfs(node.Left) || dfs(node.Right)
	}
	return dfs(root)
}
// @lc code=end

