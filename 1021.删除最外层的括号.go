/*
 * @lc app=leetcode.cn id=1021 lang=golang
 *
 * [1021] 删除最外层的括号
 */

// @lc code=start
func removeOuterParentheses(s string) string {
	var ans, st []rune
	for _, c := range s {
		if c == ')' {
			st = st[:len(st)-1]
		}
		if len(st) > 0 {
			ans = append(ans, c)
		}
		if c == '(' {
			st = append(st, c)
		}
	}
	return string(ans)
}

// @lc code=end

