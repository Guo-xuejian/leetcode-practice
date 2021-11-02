/*
 * @lc app=leetcode.cn id=14 lang=golang
 *
 * [14] 最长公共前缀
 */

// @lc code=start
func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	prefix := strs[0]
	count := len(strs)
	for i := 0; i < count; i++ {
		prefix = findCommonPart(prefix, strs[i])
		if len(prefix) == 0 {
			break
		}
	}
	return prefix
}

func findCommonPart(str1, str2 string) string {
	// 首先需要对比两个字符串长度，取较小的，避免数组越界
	// 从第一个元素开始，因为可能存在前面有共同部分，后面元素没有共同部分
	len := min(len(str1), len(str2))
	index := 0
	for index < len && str1[index] == str2[index] {
		index++
	}
	// 通过 index 截取共同部分并返回
	return str1[:index]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

