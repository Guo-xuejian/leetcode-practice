/*
 * @lc app=leetcode.cn id=429 lang=golang
 *
 * [429] N 叉树的层序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func levelOrder(root *Node) (res [][]int) {
	if root == nil {
		return
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		temp := make([]*Node, len(queue))
		copy(temp, queue)
		queue = []*Node{}
		curr := make([]int, 0, len(temp))
		for _, node := range temp {
			curr = append(curr, node.Val)
			for _, nextNode := range node.Children {
				queue = append(queue, nextNode)
			}
		}
		res = append(res, curr)
	}
	return
}

// @lc code=end

