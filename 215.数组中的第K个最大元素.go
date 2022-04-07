// 215. 数组中的第K个最大元素
// 给定整数数组 nums 和整数 k，请返回数组中第 k 个最大的元素。

// 请注意，你需要找的是数组排序后的第 k 个最大的元素，而不是第 k 个不同的元素。

// 示例 1:

// 输入: [3,2,1,5,6,4] 和 k = 2
// 输出: 5
// 示例 2:

// 输入: [3,2,3,1,2,4,5,5,6] 和 k = 4
// 输出: 4

// 提示：

// 1 <= k <= nums.length <= 104
// -104 <= nums[i] <= 104

func findKthLargest(nums []int, k int) (res int) {
	// sort.Ints(nums)
	// return nums[len(nums)-k]
	h := make(hp, len(nums))
	for idx, num := range nums {
		h[idx] = num
	}
	heap.Init(&h)
	for k > 0 {
		res = (heap.Pop(&h)).(int)
		k--
	}
	return
}

type hp []int

func (h hp) Len() int { return len(h) }

func (h hp) Less(i, j int) bool { return h[i] > h[j] }

func (h hp) Swap(i, j int) { h[i], h[j] = h[j], h[i] }

func (h *hp) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *hp) Pop() (v interface{}) {
	*h, v = (*h)[:h.Len()-1], (*h)[h.Len()-1]
	return
}