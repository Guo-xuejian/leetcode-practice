/*
 * @lc app=leetcode.cn id=589 lang=golang
 *
 * [589] N 叉树的前序遍历
 */

// @lc code=start
/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

// 递归实现
// func preorder(root *Node) (res []int) {
//     var dfs func(root *Node)
//     dfs = func(root *Node) {
//         if root != nil {
//             res = append(res, root.Val)
//             for _, node := range root.Children {
//                 dfs(node)
//             }
//         }
//     }
//     dfs(root)
//     return
// }

func preorder(root *Node) (res []int) {
	if root == nil {
		return
	}
	stack := []*Node{root}
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]                   // 实现出栈
		res = append(res, node.Val)                    // 写入结果
		for i := len(node.Children) - 1; i >= 0; i-- { // 反向写入，保证下一个节点在栈顶
			stack = append(stack, node.Children[i])
		}
	}
	return
}

// @lc code=end

