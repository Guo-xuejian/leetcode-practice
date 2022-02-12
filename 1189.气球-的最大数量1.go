/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) (res int) {
	letterMap := map[rune]int{'b': 0, 'a': 0, 'l': 0, 'o': 0, 'n': 0,}
	for _, letter := range text {
		if _, ok := letterMap[letter]; ok {
			letterMap[letter]++
		}
	}
	res = len(text)
	for key, val := range letterMap {
		if key == 'l' || key == 'o' {
			res = min(res, val / 2)
		} else {
			res = min(res, val)
		}
	}
	return
}


func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
// @lc code=end

