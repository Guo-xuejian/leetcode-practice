/*
 * @lc app=leetcode.cn id=28 lang=golang
 *
 * [28] 实现 strStr()
 */

// @lc code=start
func strStr(haystack string, needle string) int {
	length1, length2 := len(haystack), len(needle)

	for i := 0; i < length1-length2+1; i++ {
		controlStatus := false
		for j := 0; j < length2; j++ {
			if needle[j] != haystack[i+j] {
				controlStatus = true
				break
			}
		}
		if !controlStatus {
			return i
		}
	}
	return -1
}

// @lc code=end

