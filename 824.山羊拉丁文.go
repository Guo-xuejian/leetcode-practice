/*
 * @lc app=leetcode.cn id=824 lang=golang
 *
 * [824] 山羊拉丁文
 */

// @lc code=
// 元音
var Alphabet = []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}

func toGoatLatin(sentence string) string {
	sentenceSlice := strings.Fields(sentence) // 按照空格切分
	sentence = ""
	for idx, word := range sentenceSlice {
		sentence += tranfser(idx+1, word) + " "
	}
	sentence = string(sentence[:len(sentence)-1]) // 最后一个空格不要
	return sentence
}

func tranfser(index int, word string) string {
	existStatus := false
	for _, letter := range Alphabet { // 检查是否是元音开头
		if letter == word[0] {
			existStatus = true
			break
		}
	}
	// 辅音
	if !existStatus {
		word = word[1:] + string(word[0])
	}
	word += "ma"                 // 元音辅音都需要加上 ma
	for i := 0; i < index; i++ { // 按照索引添加字符 a
		word += "a"
	}
	return word
}

// @lc code=end

