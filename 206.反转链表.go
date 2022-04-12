/*
 * @lc app=leetcode.cn id=206 lang=golang
 *
 * [206] 反转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseList(head *ListNode) *ListNode {
	var dummyNode *ListNode
	for head != nil {
		temp := head.Next
		head.Next = dummyNode
		dummyNode = head
		head = temp
	}
	return dummyNode
}

// @lc code=end

