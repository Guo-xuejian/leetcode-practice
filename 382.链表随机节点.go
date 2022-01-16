/*
 * @lc app=leetcode.cn id=382 lang=golang
 *
 * [382] 链表随机节点
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
type Solution struct {
	nums  []int // 切片储存链表
	total int   // 记录总数
}

func Constructor(head *ListNode) (res Solution) {
	for node := head; node != nil; node = node.Next {
		res.nums = append(res.nums, node.Val)
		res.total++ // 计数加一
	}
	return
}

func (this *Solution) GetRandom() int {
	// 随机返回一个幸运值
	return (*this).nums[rand.Intn((*this).total)]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
// @lc code=end

