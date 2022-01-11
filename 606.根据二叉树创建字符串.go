/*
 * @lc app=leetcode.cn id=606 lang=golang
 *
 * [606] 根据二叉树创建字符串
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
func tree2str(root *TreeNode) string {
	if root == nil {
		return ""
	}
	if root.Left == nil && root.Right == nil { // 无子节点
		return strconv.Itoa(root.Val)
	}
	if root.Right == nil { // 左子节点可能存在，判断右子节点，因为该情况下如果右子节点不存在，那么右子节点对应括号缺省
		return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")"
	}
	// 都存在则递归写入
	return strconv.Itoa(root.Val) + "(" + tree2str(root.Left) + ")(" + tree2str(root.Right) + ")"
}

// @lc code=end

