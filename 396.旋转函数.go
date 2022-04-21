/*
 * @lc app=leetcode.cn id=396 lang=golang
 *
 * [396] 旋转函数
 */

// @lc code=start
func maxRotateFunction(nums []int) int {
	numSum := 0
	f := 0
	for idx, num := range nums {
		numSum += num
		f += idx * num
	}
	ans := f
	for i := len(nums) - 1; i > 0; i-- {
		f += numSum - len(nums)*nums[i]
		ans = max(ans, f)
	}
	return ans
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

