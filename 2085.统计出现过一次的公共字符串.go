/*
 * @lc app=leetcode.cn id=2085 lang=golang
 *
 * [2085] 统计出现过一次的公共字符串
 */

// @lc code=start
func countWords(words1 []string, words2 []string) (res int) {
	words1Map := make(map[string]int)
	words2Map := make(map[string]int)
	for _, one := range words1 {
		words1Map[one]++
	}
	for _, one := range words2 {
		words2Map[one]++
	}
	for key, val := range words1Map {
		if val == 1 {
			if v, ok := words2Map[key]; ok {
				if v == 1 {
					res++
				}
			}
		}
	}
	return res
}

// @lc code=end

