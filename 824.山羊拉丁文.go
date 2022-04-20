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

// 2022-04-21
// 执行用时：0 ms, 在所有 Go 提交中击败了100.00%的用户
// 内存消耗：2.2 MB, 在所有 Go 提交中击败了33.33%的用户
// func toGoatLatin(sentence string) string {
//     vowelLetterMap := map[byte]struct{}{
//         'a': struct{}{}, 'e': struct{}{}, 'i': struct{}{}, 'o': struct{}{}, 'u': struct{}{},
//         'A': struct{}{}, 'E': struct{}{}, 'I': struct{}{}, 'O': struct{}{}, 'U': struct{}{},
//     }
//     var GoatLatinTranslate func(word string) string
//     GoatLatinTranslate = func(word string) string {
//         if _, ok := vowelLetterMap[word[0]]; ok {
//             return word + "ma"
//         } else {
//             // word[0] 直接拼接时会被认作 byte
//             return word[1:] + word[0:1] + "ma"
//         }
//     }
//     sentenceSlice := strings.Split(sentence, " ")
//     for idx, word := range sentenceSlice {
//         suffix := ""
//         for i := 0; i < idx + 1; i++ {
//             suffix += "a"
//         }
//         sentenceSlice[idx] = GoatLatinTranslate(word) + suffix
//     }
//     return strings.Join(sentenceSlice, " ")
// }
// @lc code=end

