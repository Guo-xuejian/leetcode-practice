/*
 * @lc app=leetcode.cn id=234 lang=golang
 *
 * [234] 回文链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
 func isPalindrome(head *ListNode) bool {
    nodes := []int{}
    for head != nil {
        nodes = append(nodes, head.Val)
        head = head.Next
    }
    m := len(nodes)
    for i := 0; i < m / 2; i++ {
        if nodes[i] != nodes[m - i - 1] {
            return false
        }
    }
    return true
}
// @lc code=end

