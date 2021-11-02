/*
 * @lc app=leetcode.cn id=6 lang=golang
 *
 * [6] Z 字形变换
 */

// @lc code=start
func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}
	res := make([]string, numRows)
	row := 0
	flag := true
	for i := 0; i < len(s); i++ {
		if flag {
			res[row] += fmt.Sprintf("%c", s[i])
			if row+1 > numRows-1 {
				flag = false
				row--
			} else {
				row++
			}
		} else {
			res[row] += fmt.Sprintf("%c", s[i])
			if row-1 < 0 {
				flag = true
				row++
			} else {
				row--
			}
		}
	}
	resString := strings.Join(res, "")
	return resString
}

// @lc code=end

