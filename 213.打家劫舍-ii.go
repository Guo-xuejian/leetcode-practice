/*
 * @lc app=leetcode.cn id=213 lang=golang
 *
 * [213] 打家劫舍 II
 */

// @lc code=start
func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	} else if length == 2 {
		return max(nums[0], nums[1])
	}
	// 实际上和 打家劫舍 区别不大，只需要限制一下首尾相连，最简单的方式就是两个取最大，一个不给尾，一个不给头
	return max(robRange(nums[:length-1]), robRange(nums[1:]))
}

func robRange(nums []int) int {
	least, best := nums[0], max(nums[0], nums[1])
	for _, v := range nums[2:] {
		least, best = best, max(least+v, best)
	}
	return best
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

