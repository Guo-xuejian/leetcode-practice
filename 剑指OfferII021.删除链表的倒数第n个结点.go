// 剑指 Offer II 021. 删除链表的倒数第 n 个结点
// 给定一个链表，删除链表的倒数第 n 个结点，并且返回链表的头结点。

// 示例 1：

// 输入：head = [1,2,3,4,5], n = 2
// 输出：[1,2,3,5]
// 示例 2：

// 输入：head = [1], n = 1
// 输出：[]
// 示例 3：

// 输入：head = [1,2], n = 1
// 输出：[1]

// 提示：

// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	left, right := head, head
	count := 0
	for count < n { // 维持位移差
		right = right.Next
		count++
	}
	if right == nil { // 达到最大长度，删除的其实就是头节点 head
		return head.Next
	}
	for right.Next != nil { // left right 同步位移
		left = left.Next
		right = right.Next
	}
	// 此时的 left 指向倒数 n-1 个元素
	left.Next = left.Next.Next
	return head
}