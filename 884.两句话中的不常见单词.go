/*
 * @lc app=leetcode.cn id=884 lang=golang
 *
 * [884] 两句话中的不常见单词
 */

// @lc code=start
func uncommonFromSentences(s1 string, s2 string) (res []string) {
	wordTimesMap := map[string]int{}
	allWord := append(strings.Fields(s1), strings.Fields(s2)...)
	for _, word := range allWord {			// 各个单词出现次数统计
		wordTimesMap[word]++
	}

	for word, times := range wordTimesMap {	// 查找总出现次数为 1 的单词
		if times == 1 {
			res = append(res, word)
		}
	}
	return
}
// @lc code=end

