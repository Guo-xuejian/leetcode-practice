/*
 * @lc app=leetcode.cn id=678 lang=golang
 *
 * [678] 有效的括号字符串
 */

// @lc code=start
func checkValidString(s string) bool {
	var left, aster []int
	length := len(s)
	for i:=0;i<length;i++ {
		if s[i] == '(' {
			left = append(left, i)
		} else if s[i] == '*' {
			aster = append(aster, i)
		} else if len(left) > 0 {
			left = left[:len(left)-1]
		} else if len(aster) > 0 {
			aster = aster[:len(aster)-1]
		} else {
			return false
		}
	}
	i := len(left) - 1
	for j := len(aster) - 1;i>=0&&j>=0;i,j=i-1,j-1 {
		if left[i] > aster[j] {
			return false
		}
	}
	return i < 0
}
// @lc code=end

