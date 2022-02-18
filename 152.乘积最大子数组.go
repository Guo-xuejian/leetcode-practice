/*
 * @lc app=leetcode.cn id=152 lang=golang
 *
 * [152] 乘积最大子数组
 */

// @lc code=start
func maxProduct(nums []int) (res int) {
	// 由于负数的存在很有可能导致 -5 * -1 成为最大值，所以维护一个最小的
	// 反过来亦然
	maxNum, minNum, res := nums[0], nums[0], nums[0]
	for i := 1; i < len(nums); i++ {
		mx, mn := maxNum, minNum
		maxNum = max(mx*nums[i], max(nums[i], mn*nums[i]))
		minNum = min(mn*nums[i], min(nums[i], mx*nums[i]))
		res = max(maxNum, res)
	}
	return
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

