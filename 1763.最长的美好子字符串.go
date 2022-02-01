/*
 * @lc app=leetcode.cn id=1763 lang=golang
 *
 * [1763] 最长的美好子字符串
 */

// @lc code=start
func longestNiceSubstring(s string) string {
	n := len(s)
	maxPos, maxLen := 0, 0
	for i := 0; i < n; i++ {
		lower, upper := 0, 0
		for j := i; j < n; j++ {
			if s[j] >= 'a' && s[j] <= 'z' {
				// 小写字符记录进入 lower
				lower |= 1 << (int(s[j] - 'a'))
			} else {
				// 大写字符记录进入 upper
				upper |= 1 << (int(s[j] - 'A'))
			}
			// 相等意味着当前字符串中对应的小写字母都存在相应的大写字母
			if lower == upper && j - i + 1 > maxLen {
				maxPos = i			// 起点
				maxLen = j - i + 1	// 满足条件的终点
			}
		}
	}
	// 截取返回
	return string(s[maxPos:maxPos + maxLen])
}
// @lc code=end

