/*
 * @lc app=leetcode.cn id=806 lang=golang
 *
 * [806] 写字符串需要的行数
 */

// @lc code=start
func numberOfLines(widths []int, s string) []int {
	currLength, needLines := 0, 0
	for i := 0; i < len(s); i++ {
		curr := widths[s[i]-'a']
		if curr+currLength > 100 {
			needLines++
			currLength = curr
		} else {
			currLength += curr
		}
	}
	return []int{needLines + 1, currLength}
}

// @lc code=end

