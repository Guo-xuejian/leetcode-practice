/*
 * @lc app=leetcode.cn id=172 lang=golang
 *
 * [172] 阶乘后的零
 */

// @lc code=start
func trailingZeroes(n int) (res int) {
	for n != 0 {
		res, n = res+n/5, n/5
	}
	return
}

// @lc code=end

