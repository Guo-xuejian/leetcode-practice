/*
 * @lc app=leetcode.cn id=160 lang=golang
 *
 * [160] 相交链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getIntersectionNode(headA, headB *ListNode) *ListNode {
	if headA == nil || headB == nil {
		return nil
	}
	AHead, BHead := headA, headB
	for AHead != BHead {
		if AHead == nil {
			AHead = headB
		} else {
			AHead = AHead.Next
		}
		if BHead == nil {
			BHead = headA
		} else {
			BHead = BHead.Next
		}
	}
	return AHead
}

// @lc code=end

