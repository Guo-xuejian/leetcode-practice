/*
 * @lc app=leetcode.cn id=926 lang=golang
 *
 * [926] 将字符串翻转到单调递增
 */

// @lc code=start
func minFlipsMonoIncr(s string) int {
	dp0, dp1 := 0, 0
	for _, c := range s {
		dp0New, dp1New := dp0, min(dp0, dp1)
		if c == '1' {
			dp0New++
		} else {
			dp1New++
		}
		dp0, dp1 = dp0New, dp1New
	}
	return min(dp0, dp1)
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// @lc code=end

