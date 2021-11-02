/*
 * @lc app=leetcode.cn id=148 lang=golang
 *
 * [148] 排序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func sortList(head *ListNode) *ListNode {
	temp1 := head
	temp2 := temp1
	res := []int{}
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	sort.Ints(res)
	for _, one := range res {
		temp2.Val = one
		temp2 = temp2.Next
	}
	return temp1
}

// @lc code=end

