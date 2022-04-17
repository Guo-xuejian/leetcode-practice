// 请完成一个函数，输入一个二叉树，该函数输出它的镜像。

// 例如输入：

//      4
//    /   \
//   2     7
//  / \   / \
// 1   3 6   9
// 镜像输出：

//      4
//    /   \
//   7     2
//  / \   / \
// 9   6 3   1

//

// 示例 1：

// 输入：root = [4,2,7,1,3,6,9]
// 输出：[4,7,2,9,6,3,1]
//

// 限制：

// 0 <= 节点个数 <= 1000

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func mirrorTree(root *TreeNode) *TreeNode {
	if root == nil {
		return root
	}
	dummyNode := root
	queue := []*TreeNode{root}
	for len(queue) > 0 {
		temp := make([]*TreeNode, len(queue))
		copy(temp, queue)
		queue = []*TreeNode{}
		for _, node := range temp {
			node.Left, node.Right = node.Right, node.Left
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
	}
	return dummyNode
}