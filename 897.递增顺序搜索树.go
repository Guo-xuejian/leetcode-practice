/*
 * @lc app=leetcode.cn id=897 lang=golang
 *
 * [897] 递增顺序搜索树
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
func increasingBST(root *TreeNode) *TreeNode {
	allNodeVal := []int{}
	var inOrder func(root *TreeNode)
	inOrder = func(root *TreeNode) {
		// 中序遍历将所有 Val 写入 allNodeVal，由于是二叉搜索树，所以不担心有序问题
		if root != nil {
			inOrder((*root).Left)
			allNodeVal = append(allNodeVal, root.Val)
			inOrder((*root).Right)
		}
	}

	inOrder(root)

	preHead := &TreeNode{} // 辅助节点
	currNode := preHead    // 工作节点

	for _, val := range allNodeVal {
		currNode.Right = &TreeNode{Val: val} // 只有右子树
		currNode = currNode.Right
	}

	return preHead.Right
}

// @lc code=end

