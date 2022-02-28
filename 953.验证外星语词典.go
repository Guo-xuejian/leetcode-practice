/*
 * @lc app=leetcode.cn id=953 lang=golang
 *
 * [953] 验证外星语词典
 */

// @lc code=start
func isAlienSorted(words []string, order string) bool {
	letterIndexMap := map[byte]int{}
	for idx := range order {
		letterIndexMap[order[idx]] = idx
	}

	for i := 0; i < len(words)-1; i++ {
		word1, word2 := words[i], words[i+1]
		isPrefix := true

		for j := 0; j < min(len(word1), len(word2)); j++ {
			if word1[j] != word2[j] {
				if letterIndexMap[word1[j]] > letterIndexMap[word2[j]] {
					return false
				}

				isPrefix = false
				break
			}
		}

		if isPrefix && len(word1) > len(word2) {
			return false
		}
	}
	return true
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

