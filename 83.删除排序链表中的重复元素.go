/*
 * @lc app=leetcode.cn id=83 lang=golang
 *
 * [83] 删除排序链表中的重复元素
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func deleteDuplicates(head *ListNode) *ListNode {
    dummyNode := head
    preNode := head
    if head == nil {
        return head
    } else {
        head = head.Next
    }
    for head != nil {
        if head.Val == preNode.Val {
            // 存在重复元素，删除
            preNode.Next = head.Next
        } else {
            preNode = head
        }
        head = head.Next
    }
    return dummyNode
}
// @lc code=end

