// 剑指 Offer 25. 合并两个排序的链表
// 输入两个递增排序的链表，合并这两个链表并使新链表中的节点仍然是递增排序的。

// 示例1：

// 输入：1->2->4, 1->3->4
// 输出：1->1->2->3->4->4
// 限制：

// 0 <= 链表长度 <= 1000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 双指针
	if l1 == nil && l2 == nil {
		return l1
	}
	dummyNode := &ListNode{}
	pre := dummyNode
	for l1 != nil || l2 != nil {
		if l1 != nil && l2 != nil {
			if l1.Val < l2.Val {
				pre.Val = l1.Val
				l1 = l1.Next
			} else {
				pre.Val = l2.Val
				l2 = l2.Next
			}
			temp := &ListNode{}
			pre.Next = temp
			pre = temp
		} else if l1 == nil {
			pre.Val = l2.Val
			pre.Next = l2.Next
			break
		} else {
			pre.Val = l1.Val
			pre.Next = l1.Next
			break
		}
	}
	return dummyNode
}