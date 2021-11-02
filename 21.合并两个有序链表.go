/*
 * @lc app=leetcode.cn id=21 lang=golang
 *
 * [21] 合并两个有序链表
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	// 链表上递归是最好的选择
	// 首先判定两个升序链表的头节点是否为 null True 则返回另外一个链表
	// 主体内容判定的前面，就是下方代码的前两个判定，其实不仅仅是函数开始时有用，更是作为链表递归至尾部时的重要结束判定
	// 注意，判定是否为 nil 应该直接判定节点，而不是去判定节点的val和next，完全不必要的问题。。。。。
	if l1 == nil {
		return l2
	} else if l2 == nil {
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	}
}

// @lc code=end

