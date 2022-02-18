// 剑指 Offer 06. 从尾到头打印链表
// 输入一个链表的头节点，从尾到头反过来返回每个节点的值（用数组返回）。

// 示例 1：

// 输入：head = [1,3,2]
// 输出：[2,3,1]

// 限制：

// 0 <= 链表长度 <= 10000

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reversePrint(head *ListNode) (res []int) {
	for head != nil {
		res = append(res, head.Val)
		head = head.Next
	}
	// 翻转数组
	index := 0
	length := len(res)
	if length == 0 {
		return
	}
	for index <= length/2-1 {
		res[index], res[length-1-index] = res[length-1-index], res[index]
		index++
	}
	return
}