/*
 * @lc app=leetcode.cn id=58 lang=golang
 *
 * [58] 最后一个单词的长度
 */

// @lc code=start
func lengthOfLastWord(s string) (res int) {
	// 反向遍历，找到一个字符串就停下
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] != ' ' {
			res++
		} else {
			if res > 0 {
				break
			}
		}
	}
	return
}

// @lc code=end

