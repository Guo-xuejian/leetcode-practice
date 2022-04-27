// 剑指 Offer 41. 数据流中的中位数
// 如何得到一个数据流中的中位数？如果从数据流中读出奇数个数值，那么中位数就是所有数值排序之后位于中间的数值。如果从数据流中读出偶数个数值，那么中位数就是所有数值排序之后中间两个数的平均值。

// 例如，

// [2,3,4] 的中位数是 3

// [2,3] 的中位数是 (2 + 3) / 2 = 2.5

// 设计一个支持以下两种操作的数据结构：

// void addNum(int num) - 从数据流中添加一个整数到数据结构中。
// double findMedian() - 返回目前所有元素的中位数。
// 示例 1：

// 输入：
// ["MedianFinder","addNum","addNum","findMedian","addNum","findMedian"]
// [[],[1],[2],[],[3],[]]
// 输出：[null,null,null,1.50000,null,2.00000]
// 示例 2：

// 输入：
// ["MedianFinder","addNum","findMedian","addNum","findMedian"]
// [[],[2],[],[3],[]]
// 输出：[null,null,2.00000,null,2.50000]

// 限制：

// 最多会对 addNum、findMedian 进行 50000 次调用。

type IntHeap []int

func (h IntHeap) Len() int           { return len(h) }
func (h IntHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h IntHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

/** initialize your data structure here. */
type MedianFinder struct {
	large, small *IntHeap
}

/** initialize your data structure here. */
func Constructor() MedianFinder {
	a := &IntHeap{}
	b := &IntHeap{}
	heap.Init(a)
	heap.Init(b)

	return MedianFinder{a, b}

}

func (this *MedianFinder) AddNum(num int) {
	if this.large.Len() == 0 || num > (*this.large)[0] {
		heap.Push(this.large, num)
	} else {
		heap.Push(this.small, -num)
	}

	// 调节大小栈，确保 `FindMedian()` 得到正确的结果
	if this.large.Len() > this.small.Len()+1 {
		heap.Push(this.small, -heap.Pop(this.large).(int))
	} else if this.small.Len() > this.large.Len()+1 {
		heap.Push(this.large, -heap.Pop(this.small).(int))
	}
}

func (this *MedianFinder) FindMedian() float64 {
	if this.large.Len() < this.small.Len() {
		return float64(-(*this.small)[0])
	} else if this.large.Len() > this.small.Len() {
		return float64((*this.large)[0])
	}

	return float64(-(*this.small)[0]+(*this.large)[0]) / 2
}

/**
 * Your MedianFinder object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AddNum(num);
 * param_2 := obj.FindMedian();
 */