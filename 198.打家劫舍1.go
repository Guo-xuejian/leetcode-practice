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
	}
	preNum, currNum := nums[0], max(nums[0], nums[1])
	for i := 2; i < length; i++ {
		preNum, currNum = currNum, max(preNum + nums[i], currNum)
	}
	return currNum
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

