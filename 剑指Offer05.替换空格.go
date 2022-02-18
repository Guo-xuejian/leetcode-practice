// 剑指 Offer 05. 替换空格
// 请实现一个函数，把字符串 s 中的每个空格替换成"%20"。

// 示例 1：

// 输入：s = "We are happy."
// 输出："We%20are%20happy."

// 限制：

// 0 <= s 的长度 <= 10000

func replaceSpace(s string) string {
	// return strings.Replace(s, " ", "%20", -1)
	runeSlice := make([]rune, 0, len(s))
	for _, one := range s {
		if one == ' ' {
			runeSlice = append(runeSlice, []rune("%20")...)
		} else {
			runeSlice = append(runeSlice, one)
		}
	}
	return string(runeSlice)
}