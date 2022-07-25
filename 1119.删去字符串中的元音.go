// 1119. 删去字符串中的元音
// 给你一个字符串 s ，请你删去其中的所有元音字母 'a'，'e'，'i'，'o'，'u'，并返回这个新字符串。

// 示例 1：

// 输入：s = "leetcodeisacommunityforcoders"
// 输出："ltcdscmmntyfrcdrs"
// 示例 2：

// 输入：s = "aeiou"
// 输出：""

// 提示：

// 1 <= S.length <= 1000
// s 仅由小写英文字母组成
package main

func removeVowels(s string) string {
	tempSlice := []byte{}
	vowelMap := map[byte]bool{
		'a': true,
		'e': true,
		'i': true,
		'o': true,
		'u': true,
	}
	for _, letter := range s {
		if !vowelMap[byte(letter)] {
			tempSlice = append(tempSlice, byte(letter))
		}
	}
	return string(tempSlice)
}
