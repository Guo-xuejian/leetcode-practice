/*
 * @lc app=leetcode.cn id=53 lang=golang
 *
 * [53] 最大子数组和
 */

// @lc code=start
func maxSubArray(nums []int) (res int) {
	pre := 0
	res = nums[0]
	for i := 0; i < len(nums); i++ {
		pre = max(nums[i], pre+nums[i]) // 相当于维持一个单调递增数组
		res = max(res, pre)
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

