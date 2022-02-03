/*
 * @lc app=leetcode.cn id=25 lang=golang
 *
 * [25] K 个一组翻转链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func reverseKGroup(head *ListNode, k int) *ListNode {
    hair := &ListNode{Next: head}
    pre := hair

    for head != nil {
        tail := pre
        for i := 0; i < k; i++ {    // 不满足则退出
            tail = tail.Next
            if tail == nil {
                return hair.Next
            }
        }
        nex := tail.Next
        head, tail = reverse(head, tail)
        // 连接至yu原有链表
        pre.Next = head
        tail.Next = nex
        pre = tail
        head = tail.Next
    }
    return hair.Next
}

func reverse(head, tail *ListNode) (*ListNode, *ListNode) {
    // 反转链表
    prev := tail.Next
    p := head
    for prev != tail {
        nex := p.Next
        p.Next = prev
        prev = p
        p = nex
    }
    return tail, head
}
// @lc code=end

