/*
 * @lc app=leetcode.cn id=521 lang=golang
 *
 * [521] 最长特殊序列 Ⅰ
 */

// @lc code=start
func findLUSlength(a string, b string) int {
	if a == b {
		return -1
	}
	return max(len(a), len(b))
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

// @lc code=end

