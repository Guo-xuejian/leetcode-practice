/*
 * @lc app=leetcode.cn id=703 lang=golang
 *
 * [703] 数据流中的第 K 大元素
 */

// @lc code=start
type KthLargest struct {
	sort.IntSlice
	k int
}

func Constructor(k int, nums []int) KthLargest {
	kl := KthLargest{k: k}
	for _, num := range nums {
		kl.Add(num)
	}
	return kl
}

func (this *KthLargest) Push(x interface{}) {
	this.IntSlice = append(this.IntSlice, x.(int))
}

func (this *KthLargest) Pop() (v interface{}) {
	this.IntSlice, v = (this.IntSlice)[:this.Len()-1], (this.IntSlice)[this.Len()-1]
	return
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	if this.Len() > this.k {
		heap.Pop(this)
	}
	return this.IntSlice[0]
}

/**
 * Your KthLargest object will be instantiated and called as such:
 * obj := Constructor(k, nums);
 * param_1 := obj.Add(val);
 */
// @lc code=end

