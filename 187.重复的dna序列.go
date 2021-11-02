/*
 * @lc app=leetcode.cn id=187 lang=golang
 *
 * [187] 重复的DNA序列
 */

// @lc code=start
func findRepeatedDnaSequences(s string) (res []string) {
	stringMap := map[string]int{}
	lengthOfStringNeedToCount := len(s) - 10
	for i := 0; i <= lengthOfStringNeedToCount; i++ {
		sub := s[i : i+10]
		stringMap[sub]++
		if stringMap[sub] == 2 {
			res = append(res, sub)
		}
	}
	return
}

// @lc code=end

