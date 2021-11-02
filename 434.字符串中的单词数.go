/*
 * @lc app=leetcode.cn id=434 lang=golang
 *
 * [434] 字符串中的单词数
 */

// @lc code=start
func countSegments(s string) int {
	res := 0
	// 为避免重复计数，初始化 writeFlag 来表示是否仍然处在一个单词范围内
	// writeFlag 为真时放弃计数，表示还在该单词范围内
	writeFlag := false
	lengthOfString := len(s)
	for i := 0; i < lengthOfString; i++ {
		if s[i] == ' ' {
			writeFlag = false
			continue
		}
		if !writeFlag {
			res++
			writeFlag = true
		}
	}
	return res
}

// @lc code=end

