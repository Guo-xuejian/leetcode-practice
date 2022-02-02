/*
 * @lc app=leetcode.cn id=2000 lang=golang
 *
 * [2000] 反转单词前缀
 */

// @lc code=start
func reversePrefix(word string, ch byte) string {
	index := -1
	for idx, one := range word {
		if byte(one) == ch {
			index = idx
			break // 第一次出现即可
		}
	}
	if index == -1 {
		return string(word)
	}
	// 创建一个新切片,不做转化,空间大概率小一些
	temp, writeIndex, tempIndex := make([]byte, index+1), 0, index
	for tempIndex >= 0 { // 反向写入
		temp[writeIndex] = word[tempIndex]
		writeIndex++
		tempIndex--
	}
	// 转为字符串拼接
	return string(temp) + word[index+1:]
}

// @lc code=end

