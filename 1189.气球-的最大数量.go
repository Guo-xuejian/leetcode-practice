/*
 * @lc app=leetcode.cn id=1189 lang=golang
 *
 * [1189] “气球” 的最大数量
 */

// @lc code=start
func maxNumberOfBalloons(text string) int {
	letterMap := make(map[byte]int)
	for _, val := range text { // 统计 balloon 中字母在 text 出现次数
		if val == 'b' || val == 'a' || val == 'l' || val == 'o' || val == 'n' {
			letterMap[val]++
		}
	}
	// 如果字符缺失，则不可能得出答案
	if len(letterMap) != 5 {
		return 0
	}
	minNum := len(text) // 值不会超过 text 长度
	for key, val := range letterMap {
		if key == 'l' || key == 'o' { // 需要两个的
			if val/2 < minNum {
				minNum = val / 2
			}
		} else { // 只需要一个
			if val < minNum {
				minNum = val
			}
		}
	}
	// 最终返回一个短板
	return minNum
}

// @lc code=end

