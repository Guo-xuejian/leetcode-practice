/*
 * @lc app=leetcode.cn id=979 lang=golang
 *
 * [979] 在二叉树中分配硬币
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
func distributeCoins(root *TreeNode) int {
	res := 0
	dfs(root, &res)
	return res
}

func dfs(root *TreeNode, res *int) int {
	if root == nil {
		return 0
	}
	left, right := dfs(root.Left, res), dfs(root.Right, res)
	*res += abs(left) + abs(right)
	return root.Val + left + right - 1
}

func abs(x int) int {
	if x >= 0 {
		return x
	}
	return (-1) * x
}

// @lc code=end

