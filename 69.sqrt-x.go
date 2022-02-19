/*
 * @lc app=leetcode.cn id=69 lang=golang
 *
 * [69] Sqrt(x)
 */

// @lc code=start
func mySqrt(x int) (res int) {
	left, right := 0, x
	for left <= right {
		mid := (left + right) / 2
		if mid*mid <= x {
			res = mid
			left = mid + 1
		} else {
			right = mid - 1
		}
	}
	return
}

// @lc code=end

