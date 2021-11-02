/*
 * @lc app=leetcode.cn id=237 lang=golang
 *
 * [237] 删除链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode(node *ListNode) {
	// 由于题目中说明是一个有效节点且不失末尾节点，那么就不需要判定Next为空
	// 覆盖当前的节点为下一个节点的值，贪吃蛇向前走一步
	for node.Next.Next != nil { // 始终保持指针在最后一个节点之前的那一个节点
		node.Val = node.Next.Val
		node = node.Next
	}
	// 跳空最后一个，特别处理
	node.Val = node.Next.Val
	node.Next = nil // 抛弃最后一个节点
}

// @lc code=end

