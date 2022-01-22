/*
 * @lc app=leetcode.cn id=1408 lang=golang
 *
 * [1408] 数组中的字符串匹配
 */

// @lc code=start
func stringMatching(words []string) (res []string) {
	for idx1, val1 := range words {
		for idx2, val2 := range words {
			if idx1 != idx2 && InWords(val1, val2) {
				res = append(res, val1)
				break
			}
		}
	}
	return res
}

func InWords(str1, str2 string) bool {
	length1, length2 := len(str1), len(str2)
	if length1 > length2 {
		return false
	}
	for i := 0; i < length2-length1+1; i++ {
		if string(str2[i:i+length1]) == str1 {
			return true
		}
	}
	return false
}

// @lc code=end

