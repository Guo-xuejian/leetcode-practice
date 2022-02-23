/*
 * @lc app=leetcode.cn id=917 lang=golang
 *
 * [917] 仅仅反转字母
 */

// @lc code=start
func reverseOnlyLetters(s string) string {
	if len(s) == 1 {
		return s
	}
	// 转换为切片
	tempSlice := []byte(s)
	// 双指针思想，左边出现英文字符，右边找一个英文字符交换即可，注意索引不能出现实际对撞
	left, right := 0, len(s)-1
	for left < right {
		if isLetter(tempSlice[left]) { // 遇到英文字母
			// 从右往左找一个字母
			for right > left && !isLetter(tempSlice[right]) {
				right--
			}
			if right <= left { // 如果已经对撞，退出循环
				break
			}
			tempSlice[left], tempSlice[right] = tempSlice[right], tempSlice[left]
			right--
		}
		// 非英文字母不需要操作，left++ 即可
		left++
	}
	return string(tempSlice)
}

func isLetter(letter byte) bool {
	if letter >= 'a' && letter <= 'z' || letter >= 'A' && letter <= 'Z' {
		return true
	}
	return false
}

// @lc code=end

