/*
 * @lc app=leetcode.cn id=633 lang=golang
 *
 * [633] 平方数之和
 */

// @lc code=start
func judgeSquareSum(c int) bool {
	left, right := 0, int(math.Sqrt(float64(c)))
	for left <= right {
		sum := left*left + right*right
		if sum == c {
			return true
		} else if sum > c {
			right--
		} else {
			left++
		}
	}
	return false
}

// @lc code=end

