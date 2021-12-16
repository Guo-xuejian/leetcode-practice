/*
 * @lc app=leetcode.cn id=198 lang=golang
 *
 * [198] 打家劫舍
 */

// @lc code=start
func rob(nums []int) int {
	length := len(nums)
	if length == 1 {
		return nums[0]
	} else if length == 2 {
		return max(nums[0], nums[1])
	}

	// 第 k 个最大值判定，如果 k 和 k-2 更大则取最大值，否则该位的最大值判定为 k-1 与之前组合
	least, best := nums[0], max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		least, best = best, max(least+nums[i], best)
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

