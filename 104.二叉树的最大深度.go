/*
 * @lc app=leetcode.cn id=104 lang=golang
 *
 * [104] 二叉树的最大深度
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
package main

func maxDepth(root *TreeNode) int {
	// 从两个分支中选出一个最大的，对于子节点也是如此，递归调用即可
	if root != nil {
		leftLength := maxDepth(root.Left)
		rightLength := maxDepth(root.Right)
		return max(leftLength, rightLength) + 1 // 首节点需要额外加上
	} else {
		return 0
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

