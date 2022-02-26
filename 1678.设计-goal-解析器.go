/*
 * @lc app=leetcode.cn id=1678 lang=golang
 *
 * [1678] 设计 Goal 解析器
 */

// @lc code=start
func interpret(command string) string {
	// return strings.Replace(strings.Replace(command, "()", "o", -1), "(al)", "al", -1)
	preLetter := command[0]
	index := 1
	res := make([]byte, 0, len(command))
	if preLetter == 'G' {
		res = append(res, 'G')
	}
	for index < len(command) {
		if command[index] == ')' && preLetter == '(' {
			res = append(res, 'o')
		} else if command[index] == 'a' {
			preLetter = ')'
			index += 3
			res = append(res, []byte("al")...)
			continue
		} else if command[index] == 'G' {
			res = append(res, 'G')
		}
		preLetter = command[index]
		index++
	}
	return string(res[:len(res)])
}

// @lc code=end

