/*
 * @lc app=leetcode.cn id=461 lang=golang
 *
 * [461] 汉明距离
 */

// @lc code=start
func hammingDistance(x int, y int) (res int) {
	for s := x ^ y; s > 0; s >>= 1 {
		res += s & 1
	}
	return
}

// @lc code=end

