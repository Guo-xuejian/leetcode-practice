/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	/*
		head := &ListNode{}
		temp := head
		sumOfAll := 0
		n1 := 0
		n2 := 0
		for l1 != nil || l2 != nil || sumOfAll > 0 {
			n1 = 0
			n2 = 0
			// go 不支持三元表达式,所以必须得有两个变量记录两个链表的数字
			if l1 != nil {
				n1 = l1.Val
				l1 = l1.Next
			}
			if l2 != nil {
				n2 = l2.Val
				l2 = l2.Next
			}
			sumOfAll += n1 + n2
			temp.Next = &ListNode{Val:sumOfAll % 10}
			temp = temp.Next
			sumOfAll = sumOfAll / 10
		}
		return head.Next
	*/
	// 由于两个链表顺序颠倒，所以存在的位数是一一对应的，可以考虑对两位相加，超过10的部分继续传递————作为其中一个判定条件
	head := &ListNode{Val: -1} // 由于都是>= 0，可以给一个负数用来作为起点
	result := &ListNode{}
	restPart := 0
	n1 := 0
	n2 := 0
	sum := 0
	for l1 != nil || l2 != nil || restPart > 0 {
		n1 = 0
		n2 = 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum = n1 + n2 + restPart
		restPart = sum / 10
		fmt.Println(restPart)
		sum = sum % 10
		// 节点 Val < 0 ，那么就是起点
		if head.Val < 0 {
			head.Val = sum
			result = head
		} else {
			result.Next = &ListNode{Val: sum}
			result = result.Next
		}
	}
	return head
}

// @lc code=end

