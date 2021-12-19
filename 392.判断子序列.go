/*
 * @lc app=leetcode.cn id=392 lang=golang
 *
 * [392] 判断子序列
 */

// @lc code=start
func isSubsequence(s string, t string) bool {
	if s == "" {
		return true
	} else if s != "" && t == "" {
		return false
	}
	// 维护当前字母和上一个字母的索引
	preLetterIndex := -1
	afterLetterIndex := -1
	for i := 0; i < len(s); i++ {
		existStatus := false
		for j := 0; j < len(t); j++ {
			if s[i] == t[j] {
				existStatus = true
				if j > afterLetterIndex { // 只要出现一个满足条件的即可
					afterLetterIndex = j
					break
				}
			}
		}
		// 存在且索引符合条件则重置索引，否则直接返回 false 因为不存在或者索引不满足条件
		if existStatus && afterLetterIndex > preLetterIndex {
			preLetterIndex = afterLetterIndex
		} else {
			return false
		}
	}
	return true
}

// @lc code=end

