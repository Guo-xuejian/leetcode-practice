/*
 * @lc app=leetcode.cn id=2104 lang=golang
 *
 * [2104] 子数组范围和
 */

// @lc code=start
func subArrayRanges(nums []int) (res int64) {
	m := len(nums)
	for i := 0; i < m; i++ {
		minVal, maxVal := math.MaxInt64, -math.MaxInt64
		for j := i; j < m; j++ {
			minVal = min(nums[j], minVal)
			maxVal = max(nums[j], maxVal)
			res += int64(maxVal) - int64(minVal)
		}
	}
	return res
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

