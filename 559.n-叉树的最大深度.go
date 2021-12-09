/*
 * @lc app=leetcode.cn id=559 lang=golang
 *
 * [559] N 叉树的最大深度
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func maxDepth(root *Node) int {
	if root == nil { // 一个空节点意味着上一节点为叶子节点，此节点实际上是虚构的一个节点
		return 0
	}
	maxChildDepth := 0
	for _, child := range root.Children {
		if childDepth := maxDepth(child); childDepth > maxChildDepth {
			maxChildDepth = childDepth // 取子节点中深度较大者
		}
	}
	return maxChildDepth + 1 // maxChildDepth 对应子节点大的最大深度，加上本身节点的高度 1
}

// @lc code=end

