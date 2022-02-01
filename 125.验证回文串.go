/*
 * @lc app=leetcode.cn id=125 lang=golang
 *
 * [125] 验证回文串
 */

// @lc code=start
func isPalindrome(s string) bool {
	left, right := 0, len(s) - 1
	for left < right {
		numStatus := false
		// 越界判定要在前面
		for left < right && !isValid(s[left]) {
			// 不为有效字符左指针前移
			left++
		}
		for left < right && !isValid(s[right]) {
			// 同上,右指针右移
			right--
		}
		if isNumeric(s[left]) || isNumeric(s[right]) {
			// 存在数字
			numStatus = true
		}
		// 字符相同,字母大小写判定，大小写不能出现数字
		if s[left] == s[right] || (!numStatus && (int(s[left]) + 32 == int(s[right]) || int(s[right]) + 32 == int(s[left]))) {
			left++
			right--
		} else {
			return false
		}
	}
	return true
}

func isValid(s byte) bool {
	// 有效字符为数字,大小写英文字母
	if (s >= '0' && s <= '9') || (s >= 'a' && s <= 'z') || (s >= 'A' && s <= 'Z') {
		return true
	}
	return false
}

func isNumeric(s byte) bool {
	if s >= '0' && s <= '9' {
		return true
	}
	return false
}
// @lc code=end

