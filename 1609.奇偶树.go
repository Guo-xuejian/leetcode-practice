/*
 * @lc app=leetcode.cn id=1609 lang=golang
 *
 * [1609] 奇偶树
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
func isEvenOddTree(root *TreeNode) bool {
	queue := []*TreeNode{root}
	level := 0
	maxNum := math.MaxInt32
	for len(queue) > 0 {
		prev := 0
		if level%2 == 1 {
			prev = maxNum
		}
		next := []*TreeNode{}
		for _, node := range queue {
			val := node.Val
			// 偶数层不出现偶数，奇数层不出现奇数
			if val%2 == level%2 {
				return false
			}
			// 偶数层单调递增，局部递减则不符合要求
			if level%2 == 0 && val <= prev {
				return false
			}
			// 奇数层单调递减，局部递增则不符合要求
			if level%2 == 1 && val >= prev {
				return false
			}
			prev = val
			if node.Left != nil {
				next = append(next, node.Left)
			}
			if node.Right != nil {
				next = append(next, node.Right)
			}
		}
		// 覆盖当前层，叶子节点层时 next 为空则退出循环
		queue = next
		level++ // 去到下一层
	}
	return true
}

// @lc code=end

