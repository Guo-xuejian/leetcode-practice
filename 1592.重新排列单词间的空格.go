/*
 * @lc app=leetcode.cn id=1592 lang=golang
 *
 * [1592] 重新排列单词间的空格
 */

// @lc code=start
func reorderSpaces(text string) string {
	spaceNum := 0           // 空格数
	wordSlice := []string{} // 单词切片
	word := ""
	for i := 0; i < len(text); i++ {
		if text[i] >= 97 && text[i] <= 122 {
			word += string(text[i]) // 字符写入
		} else {
			spaceNum++         // 非字符添加空格
			if len(word) > 0 { // 出现空格时判定单词情况并写入切片
				wordSlice = append(wordSlice, word)
				word = "" // 重置以待下一个单词
			}
		}
	}
	// 结尾处单词后没有空格处理
	if len(word) > 0 {
		wordSlice = append(wordSlice, word)
	}
	spaceIntervalNum := len(wordSlice) - 1 // 间隔数
	spaceInterval := ""                    // 最终间隔处写入空格字符串
	spaceLeft := ""                        // 最后结尾处空格字符串
	if spaceIntervalNum == 0 {             // 只有一个单词，无间隔，存在的空格全部写如结尾
		spaceInterval = ""
		for i := 0; i < spaceNum; i++ {
			spaceLeft += " "
		}
	} else { // 存在间隔
		for i := 0; i < spaceNum/spaceIntervalNum; i++ {
			spaceInterval += " "
		}
		for i := 0; i < spaceNum%spaceIntervalNum; i++ {
			spaceLeft += " "
		}
	}
	text = ""
	for idx, one := range wordSlice {
		text += one
		if idx == spaceIntervalNum { // 结束索引就是间隔数
			text += spaceLeft // 结束时写入结尾空格字符串
			break
		}
		// 未结束写入间隔字符串
		text += spaceInterval
	}
	return text
}

// @lc code=end

