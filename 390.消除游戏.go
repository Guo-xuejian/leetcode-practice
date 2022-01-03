/*
 * @lc app=leetcode.cn id=390 lang=golang
 *
 * [390] 消除游戏
 */

// @lc code=start
func lastRemaining(n int) int {
	res := 1
	step := 1
	left := true

	for n > 1 {
		// 左删或者右删时为奇数情况下都会删除第一个数字
		if left || n%2 != 0 {
			res += step
		}
		step <<= 1   // 步长变为两倍
		n >>= 1      // n 缩减为一半
		left = !left // 方向改变
	}
	return res
}

// @lc code=end

