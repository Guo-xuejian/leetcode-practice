/*
 * @lc app=leetcode.cn id=515 lang=golang
 *
 * [515] 在每个树行中找最大值
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
func largestValues(root *TreeNode) (res []int) {
	if root == nil {
		return
	}
	queue := []*TreeNode{root}
	for len(queue) != 0 {
		temp, maxNum := []*TreeNode{}, queue[0].Val
		for _, node := range queue {
			if node.Val > maxNum {
				maxNum = node.Val
			}
			if node.Left != nil {
				temp = append(temp, node.Left)
			}
			if node.Right != nil {
				temp = append(temp, node.Right)
			}
		}
		res = append(res, maxNum)
		// queue = temp[:]
		queue = make([]*TreeNode, len(temp))
		copy(queue, temp)
	}
	return
}

// @lc code=end

