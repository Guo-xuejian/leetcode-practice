/*
 * @lc app=leetcode.cn id=102 lang=golang
 *
 * [102] 二叉树的层序遍历
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

func levelOrder(root *TreeNode) (res [][]int) {
	if root == nil {
		return
	}
	// 难点在于如何记住每一层======> 把每一层都写入队列，遍历时将其子树写入新队列后覆盖
	q := []*TreeNode{root} // 第一层有一个根节点
	for i := 0; len(q) > 0; i++ {
		res = append(res, []int{})
		p := []*TreeNode{}            // 即将替代 p 的切片
		for j := 0; j < len(q); j++ { // 根据当前层所拥有的节点数量
			node := q[j]
			res[i] = append(res[i], node.Val)
			if node.Left != nil { // 写入左子树
				p = append(p, node.Left)
			}
			if node.Right != nil { // 写入右子树
				p = append(p, node.Right)
			}
		}
		q = p // 迭代至下一层
	}
	return
}

// @lc code=end

