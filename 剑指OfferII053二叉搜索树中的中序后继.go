/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func inorderSuccessor(root *TreeNode, p *TreeNode) *TreeNode {
	res := &TreeNode{}
	res = nil // 需要防止不存在的情况，该情况下 Val 会有默认值 0
	for root != nil {
		if root.Val > p.Val {
			res = root
			root = root.Left
		} else {
			root = root.Right
		}
	}
	return res
}