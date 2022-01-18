/*
 * @lc app=leetcode.cn id=1790 lang=golang
 *
 * [1790] 仅执行一次字符串交换能否使两个字符串相等
 */

// @lc code=start
func areAlmostEqual(s1 string, s2 string) bool {
	res := []int{}
	for i := 0; i < len(s1); i++ {
		if s1[i] != s2[i] {
			res = append(res, i) // 写入不同字符索引
		}
		if len(res) > 2 {
			return false
		}
	}
	if len(res) == 1 {
		return false
	}
	// 长度为 0 字符串无不同 或者 长度为 2 差异处正好互补
	if len(res) == 0 || s1[res[0]] == s2[res[1]] && s1[res[1]] == s2[res[0]] {
		return true
	}
	return false
}

// @lc code=end

