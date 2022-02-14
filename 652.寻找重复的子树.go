/*
 * @lc app=leetcode.cn id=652 lang=golang
 *
 * [652] 寻找重复的子树
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

func findDuplicateSubtrees(root *TreeNode) (res []*TreeNode) {
	allTreeMap := map[string]int{}

	// 深度优先，过程中序列化所有的子树并存储次数于 allTreeMap，达到 2 写入 res
	var postOrder func(root *TreeNode) string
	postOrder = func(root *TreeNode) string {
		if root == nil {
			return "#"
		}

		left := postOrder(root.Left)
		right := postOrder(root.Right)

		subTree := left + "," + right + "," + strconv.Itoa(root.Val)
		
		// 不需要判定直接加，加完后判定下是不是 2 就可以，只能写入 res 1 次
		allTreeMap[subTree]++
		if allTreeMap[subTree] == 2 {
			res = append(res, root)
		}
		return subTree
	}

	if root != nil {	// 空不需要继续
		postOrder(root)
	}
	return
}
// @lc code=end

