// 面试题 02.06. 回文链表
// 编写一个函数，检查输入的链表是否是回文的。

// 示例 1：

// 输入： 1->2
// 输出： false
// 示例 2：

// 输入： 1->2->2->1
// 输出： true

// 进阶：
// 你能否用 O(n) 时间复杂度和 O(1) 空间复杂度解决此题？

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	allValues := []int{}
	for head != nil {
		allValues = append(allValues, head.Val) // 逐个写入
		head = head.Next
	}
	length := len(allValues)
	for idx, value := range allValues[:length/2] {
		if value != allValues[length-1-idx] { // 对称位比较，不同则不是回文
			return false
		}
	}
	return true
}