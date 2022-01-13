/*
 * @lc app=leetcode.cn id=2124 lang=golang
 *
 * [2124] 检查是否所有 A 都在 B 之前
 */

// @lc code=start
func checkString(s string) bool {
	aStatus := false
	bStatus := false
	a := byte('a')
	// 遍历，如果 a 出现时 b 已经出现（即 bStatus 为 true），则返回 false
	for i := 0; i < len(s); i++ {
		if s[i] == a {
			if !aStatus {
				aStatus = true
			}
			if bStatus {
				return false
			}
		} else {
			if !bStatus {
				bStatus = true
			}
		}
	}
	return true

}

// @lc code=end

