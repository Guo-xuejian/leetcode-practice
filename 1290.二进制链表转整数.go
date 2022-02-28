/*
 * @lc app=leetcode.cn id=1290 lang=golang
 *
 * [1290] 二进制链表转整数
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getDecimalValue(head *ListNode) (res int) {
	// 题目说明不为空，则不需要判定
	for head != nil {
		res = res*2 + head.Val
		head = head.Next
	}
	return
}

// @lc code=end

