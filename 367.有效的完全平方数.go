/*
 * @lc app=leetcode.cn id=367 lang=golang
 *
 * [367] 有效的完全平方数
 */

// @lc code=start
func isPerfectSquare(num int) bool {
	left, right := 0, num
	for left <= right {
		mid := (right-left)>>2 + left
		curr := mid * mid
		if curr == num {
			return true
		} else if curr > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}

// @lc code=end

