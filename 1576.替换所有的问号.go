/*
 * @lc app=leetcode.cn id=1576 lang=golang
 *
 * [1576] 替换所有的问号
 */

// @lc code=start
func modifyString(s string) (res string) {
	chars := make([]byte, len(s))
	for i := 0; i < len(chars); i++ {
		if s[i] != '?' {
			chars[i] = s[i]
			continue
		}
		for j := byte('a'); j < 'z'; j++ {
			// ？替代元素和前者后者皆不相同
			if (i == 0 || j != chars[i-1]) && (i == len(chars)-1 || j != s[i+1]) {
				chars[i] = j
				break
			}
		}
	}
	return string(chars)
}

// @lc code=end

