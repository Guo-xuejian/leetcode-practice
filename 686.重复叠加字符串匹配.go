/*
 * @lc app=leetcode.cn id=686 lang=golang
 *
 * [686] 重复叠加字符串匹配
 */

// @lc code=start
func repeatedStringMatch(a string, b string) int {
	res := 1
	// 直接包含直接返回
	if strings.Contains(a, b) {
		return res
	}
	// 字符必须要一致
	for i := 0; i < len(b); i++ {
		if strings.Index(a, string(b[i])) == -1 {
			return -1
		}
	}
	ALength := len(a)
	maxLength := 2*ALength + len(b)
	tempString := a
	tempLength := ALength
	for tempLength <= maxLength {
		res++
		tempString += a
		tempLength += ALength
		if strings.Index(tempString, b) != -1 {
			return res
		}
	}
	return -1

}

// @lc code=end

