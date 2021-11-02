/*
 * @lc app=leetcode.cn id=24 lang=golang
 *
 * [24] 两两交换链表中的节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func swapPairs(head *ListNode) *ListNode {
	dummyHead := &ListNode{0, head} // 结构体直接指定位置給值，map 才需要指定键名
	temp := dummyHead
	for temp.Next != nil && temp.Next.Next != nil {
		node1 := temp.Next      // node1 指向一号元素
		node2 := temp.Next.Next // node2 指向二号元素
		temp.Next = node2       // 零号节点指向二号节点
		// 断开一号节点和二号节点
		node1.Next = node2.Next // 一号节点指向二号节点的下一节点
		node2.Next = node1      // 二号节点指向一号节点
		// 完成交换，temp 指向二号元素，跳跃一位是因为题目要求两两交换，不是每一位都交换（12，34而不是23交换）
		temp = node1 // node1 正好满足要求
	}
	return dummyHead.Next
}

// @lc code=end

