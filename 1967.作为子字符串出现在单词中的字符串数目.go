/*
 * @lc app=leetcode.cn id=1967 lang=golang
 *
 * [1967] 作为子字符串出现在单词中的字符串数目
 */

// @lc code=start
func numOfStrings(patterns []string, word string) (res int) {
	for _, one := range patterns {
		if find(one, word) {
			res++
		}
	}
	return
}

func find(one, word string) bool {
	oneLength, wordLength := len(one), len(word)
	if oneLength > wordLength { // 长度不能大于
		return false
	}
	for i := 0; i < wordLength-oneLength+1; i++ {
		if string(word[i:i+oneLength]) == one { // 找到
			return true
		}
	}
	return false // 没找到
}

// @lc code=end

