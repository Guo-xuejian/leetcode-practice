/*
 * @lc app=leetcode.cn id=230 lang=golang
 *
 * [230] 二叉搜索树中第K小的元素
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
func kthSmallest(root *TreeNode, k int) int {
	res := []int{}     // 定义切片储存数据
	resPointer := &res // 传递切片的指针
	insertIntoRes(root, resPointer)
	sort.Ints(res)  // 数组排序 小==》大
	return res[k-1] // 由于要求从 1 开始，返回 k-1 即可
}

func insertIntoRes(root *TreeNode, resPointer *[]int) {
	if root != nil { // 当前节点不为空则递归调用
		(*resPointer) = append((*resPointer), root.Val)
		insertIntoRes(root.Left, resPointer)
		insertIntoRes(root.Right, resPointer)
	}
	return
}

// @lc code=end

