/*
 * @lc app=leetcode.cn id=543 lang=golang
 *
 * [543] 二叉树的直径
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
 func diameterOfBinaryTree(root *TreeNode) int {
    res := 1
    var depth func(node *TreeNode) int
    depth = func(node *TreeNode) int {
        if node == nil {
            return 0
        }
        L, R := depth(node.Left), depth(node.Right)
        res = max(res, L + R + 1)
        return max(R, L) + 1
    }
    depth(root)
    return res - 1
}

func max(x, y int) int {
    if x > y {
        return x
    }
    return y
}
// @lc code=end

