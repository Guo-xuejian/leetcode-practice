/*
 * @lc app=leetcode.cn id=1768 lang=golang
 *
 * [1768] 交替合并字符串
 */

// @lc code=start
func mergeAlternately(word1 string, word2 string) string {
	index1, index2 := 0, 0
	m, n := len(word1), len(word2)
	writeIndex := 0
	res := make([]byte, m+n)
	for index1 < m || index2 < n {
		if index1 < m {
			res[writeIndex] = word1[index1]
			index1++
			writeIndex++
		}
		if index2 < n {
			res[writeIndex] = word2[index2]
			index2++
			writeIndex++
		}
	}
	return string(res)
}

// @lc code=end

