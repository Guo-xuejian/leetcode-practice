/*
 * @lc app=leetcode.cn id=796 lang=golang
 *
 * [796] 旋转字符串
 */

// @lc code=start
func rotateString(s string, goal string) bool {
	m, n := len(s), len(goal)
	if m > n {
		return false
	}
	s = s + s
	m *= 2
	for i := 0; i < m; i++ {
		if s[i] == goal[0] {
			index1 := i
			index2 := 0
			for index1 < m && index2 < n && s[index1] == goal[index2] {
				index1++
				index2++
			}
			if index2 == n {
				return true
			}
		}
	}
	return false
}

// @lc code=end

