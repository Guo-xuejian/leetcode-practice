/*
 * @lc app=leetcode.cn id=1614 lang=golang
 *
 * [1614] 括号的最大嵌套深度
 */

// @lc code=start
func maxDepth(s string) (res int) {
	left := 0
	for _, one := range s {
		if one == '(' {
			left++
			if left > res {
				res = left
			}
		} else if one == ')' {
			if left > 0 {
				left--
			}
		}
	}
	return res
}

// @lc code=end

