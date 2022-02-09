/*
 * @lc app=leetcode.cn id=942 lang=golang
 *
 * [942] 增减字符串匹配
 */

// @lc code=start
func diStringMatch(s string) (res []int) {
	minNum, maxNum := 0, len(s)
	for _, letter := range s {
		if letter == 'I' {
			res = append(res, minNum)
			minNum++
		} else {
			res = append(res, maxNum)
			maxNum--
		}
	}
	res = append(res, minNum)
	return
}
// @lc code=end

