/*
 * @lc app=leetcode.cn id=384 lang=golang
 *
 * [384] 打乱数组
 */

// @lc code=start
type Solution struct {
	nums, original []int
}

func Constructor(nums []int) Solution {
	// 引用传递需要做一个拷贝
	return Solution{nums, append([]int{}, nums...)}
}

func (this *Solution) Reset() []int {
	copy(this.nums, this.original) // 做拷贝，不直接 =
	return this.nums
}

func (this *Solution) Shuffle() []int {
	shuffledList := make([]int, len(this.nums))
	for i := range shuffledList {
		j := rand.Intn(len(this.nums))
		shuffledList[i] = this.nums[j]
		// 避免重复写入，删除该元素，最后覆盖数组即可
		this.nums = append(this.nums[:j], this.nums[j+1:]...)
	}
	this.nums = shuffledList // 覆盖数组
	return this.nums
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.Reset();
 * param_2 := obj.Shuffle();
 */
// @lc code=end

