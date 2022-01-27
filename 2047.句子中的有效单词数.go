/*
 * @lc app=leetcode.cn id=2047 lang=golang
 *
 * [2047] 句子中的有效单词数
 */

// @lc code=start
func countValidWords(sentence string) (res int) {
	for _, val := range strings.Fields(sentence) { // 按空格切分遍历
		if valid(val) {
			res++ // 满足加一
		}
	}
	return
}

func valid(str string) bool {
	connectStatus := false // 记录当前字符串中 - 是否出现
	length := len(str)
	for i, val := range str {
		if val >= '0' && val <= '9' { // token 不含数字
			return false
		}
		if (val == '!' || val == '.' || val == ',') && i < length-1 { // 连接符不能在结尾之前出现
			return false
		}

		if val == '-' {
			// 连接符至多一个，True 则错误，出现在首尾错误，连字符两侧应当都存在小写字母
			if connectStatus || i == 0 || i == length-1 || !(str[i-1] >= 'a' && str[i-1] <= 'z') || !(str[i+1] >= 'a' && str[i+1] <= 'z') {
				return false
			}
			connectStatus = true // 首尾之外能够出现一次
		}
	}
	return true // 无异常则为 token
}

// @lc code=end

