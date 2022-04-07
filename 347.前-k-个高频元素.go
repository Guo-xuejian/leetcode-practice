/*
 * @lc app=leetcode.cn id=347 lang=golang
 *
 * [347] 前 K 个高频元素
 */

// @lc code=start
func topKFrequent(nums []int, k int) []int {
	numMap := map[int]int{}
	for _, num := range nums {
		numMap[num]++
	}
	h := &hp{}
	heap.Init(h)
	for key, value := range numMap {
		heap.Push(h, [2]int{key, value})
		if h.Len() > k {
			heap.Pop(h)
		}
	}
	ret := make([]int, k)
	for i := 0; i < k; i++ {
		ret[k-i-1] = heap.Pop(h).([2]int)[0]
	}
	return ret
}

type hp [][2]int

func (h hp) Len() int           { return len(h) }
func (h hp) Less(i, j int) bool { return h[i][1] < h[j][1] }
func (h hp) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.([2]int))
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}

// @lc code=end

