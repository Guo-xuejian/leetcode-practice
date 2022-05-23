/*
 * @lc app=leetcode.cn id=965 lang=golang
 *
 * [965] 单值二叉树
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
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.1 MB, 在所有 Go 提交中击败了45.83%的用户
func isUnivalTree(root *TreeNode) bool {
    preNum := root.Val  // 记录根节点的值，后续不相同直接返回
    var dfs func(node *TreeNode) bool
    dfs = func(node *TreeNode) bool {
        if node == nil {    // 正常结束
            return false
        } else if node.Val != preNum {  // 不相等返回
            return true
        }
        // 判定左右子节点
        return dfs(node.Left) || dfs(node.Right)
    }
    // 根节点的左右子节点判定开始即可
    if dfs(root.Left) || dfs(root.Right) {  // 其中有一个不相同
        return false
    }
    // 满足单值二叉树定义
    return true
}
// @lc code=end

