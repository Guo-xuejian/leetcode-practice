/*
 * @lc app=leetcode.cn id=709 lang=golang
 *
 * [709] 转换成小写字母
 */

// @lc code=start
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：1.9 MB, 在所有 Go 提交中击败了98.64%的用户
func toLowerCase(s string) string {
	res := make([]rune, len(s))
	for idx, letter := range s {
		if letter >= 'A' && letter <= 'Z' {
			res[idx] = rune(int(letter) ^ 32)
		} else {
			res[idx] = letter
		}
	}
	return string(res)
}

// @lc code=end

