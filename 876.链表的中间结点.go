/*
 * @lc app=leetcode.cn id=876 lang=golang
 *
 * [876] 链表的中间结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func middleNode(head *ListNode) *ListNode {
	// dummyNode := head
	// nodeCount := 0
	// for head != nil {
	//     nodeCount++
	//     head = head.Next
	// }
	// nodeCount /= 2
	// for nodeCount != 0 {
	//     dummyNode = dummyNode.Next
	//     nodeCount--
	// }
	// return dummyNode

	// 快慢指针
	slow := head
	for head != nil && head.Next != nil {
		slow = slow.Next
		head = head.Next.Next
	}
	return slow
}

// @lc code=end

