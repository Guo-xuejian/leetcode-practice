// 剑指 Offer 50. 第一个只出现一次的字符
// 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

// 示例 1:

// 输入：s = "abaccdeff"
// 输出：'b'
// 示例 2:

// 输入：s = ""
// 输出：' '

// 限制：

// 0 <= s 的长度 <= 50000

func firstUniqChar(s string) byte {
	if len(s) == 0 {
		return ' '
	}
	letterMap := map[byte]int{}
	for i := 0; i < len(s); i++ {
		letterMap[s[i]]++
	}
	// 由于 Map 的遍历是无序的，所以需要借助原字符串 s
	// for key, val := range letterMap {
	//     if val == 1 {
	//         return key
	//     }
	// }
	for i := 0; i < len(s); i++ {
		if letterMap[s[i]] == 1 {
			return s[i]
		}
	}
	return ' '
}