/*
 * @lc app=leetcode.cn id=357 lang=golang
 *
 * [357] 计算各个位数不同的数字个数
 */

// @lc code=start
func countNumbersWithUniqueDigits(n int) int {
	if n == 0 {
		return 1
	}
	if n == 1 {
		return 10
	}
	res, curr := 10, 9.
	for i := 0; i < n-1; i++ {
		curr *= 9 - i
		res += curr
	}
	return res
}

// @lc code=end

