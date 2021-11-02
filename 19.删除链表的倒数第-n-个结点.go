/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func getLength(head *ListNode) (length int) {
	for ; head != nil; head = head.Next {
		length++
	}
	return
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	length := getLength(head)
	temp := &ListNode{0, head}      // 返回节点的头节点
	cur := temp                     // 实际操作链表
	for i := 0; i < length-n; i++ { // length-n 不能到达，因为实际操作的节点是删除节点的上一个节点
		cur = cur.Next
	}
	cur.Next = cur.Next.Next
	return temp.Next
}

// @lc code=end

