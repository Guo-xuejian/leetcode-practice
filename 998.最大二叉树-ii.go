/*
 * @lc app=leetcode.cn id=998 lang=golang
 *
 * [998] 最大二叉树 II
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
func insertIntoMaxTree(root *TreeNode, val int) *TreeNode {
	var parent *TreeNode
	for cur := root; cur != nil; cur = cur.Right {
		if val > cur.Val {
			if parent == nil {
				return &TreeNode{val, root, nil}
			}
			parent.Right = &TreeNode{val, cur, nil}
			return root
		}
		parent = cur
	}
	parent.Right = &TreeNode{Val: val}
	return root
}

// @lc code=end

