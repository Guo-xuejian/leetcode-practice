/*
 * @lc app=leetcode.cn id=107 lang=golang
 *
 * [107] 二叉树的层序遍历 II
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
func levelOrderBottom(root *TreeNode) (res [][]int) {
	_ = dfs(root, 0, &res)
	// 翻转切片
	length := len(res)
	for i := 0; i < length / 2; i++ {
		res[i], res[length - 1 - i] = res[length - 1 - i], res[i]
	}
	return res
}

func dfs(root *TreeNode, depth int, res *[][]int) []int {
	if root == nil {	// 叶子节点之后不再继续
		return []int{}
	}
	// 由于递归，可能会导致 len(*res) != depth，判定是否需要添加下一层节点的切片
	if len(*res) == depth {
		*res = append(*res, []int{})
	}
	// 当前节点值加入
	(*res)[depth] = append((*res)[depth], root.Val)

	if root.Left != nil {	// 左子节点
		dfs(root.Left, depth + 1, res)
	}

	if root.Right != nil {	// 右子节点
		dfs(root.Right, depth + 1, res)
	}
	return []int{}	// 必须返回一个，主函数忽略即可
}
// @lc code=end

