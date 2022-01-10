/*
 * @lc app=leetcode.cn id=2095 lang=golang
 *
 * [2095] 删除链表的中间节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteMiddle(head *ListNode) *ListNode {
	if head.Next == nil { // 单节点只能删除该节点，最后为 空节点
		return &ListNode{}
	}
	// 快慢指针，由于中间节点，而slow和fast一起走的话只能满足slow到中间节点，需要一个pre_slow来绑定slow前一个位置，最终也就是中间节点的前一个节点，
	slow, fast, slowPre := head, head, &ListNode{}
	for fast != nil && fast.Next != nil { // 由于 fast 一次走两个节点，所以 fast 和 fast.Next 都需要判定，避免出现空指针 panic
		fast = fast.Next.Next // 一次走两个节点，倍数差 2
		slowPre = slow        // 记录当前中间节点的前一个节点
		slow = slow.Next      // 一次一个节点，最终就是中间节点
	}
	slowPre.Next = slowPre.Next.Next // 中间节点删除
	return head                      // 返回头节点
}

// @lc code=end

