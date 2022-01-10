/*
 * @lc app=leetcode.cn id=306 lang=golang
 *
 * [306] 累加数
 */

// @lc code=start
func isAdditiveNumber(num string) bool {
	for i := 0; i < len(num)-1; i++ { // 遍历第一个数字的边界
		// 首位可以为 0，但是 0 必须要单独一个数字, i > 1 而不是 > 0 是因为截取操作的右边界是不包含的，所以需要等于 1 才能拿到至少一个字符，后续就是大于了
		if num[0] == '0' && i > 1 {
			return false
		}
		for j := i + 1; j < len(num); j++ { // 遍历第二个数字的边界
			if j-1 > 1 && num[i] == '0' {
				break // 前导符不能为 0
			}
			if dfs(num[:j], string(num[:i]), string(num[i:j])) {
				return true
			}
		}
	}
	return false
}

func dfs(num string, first, second string) bool {
	if len(num) == 0 {
		return true // 空满足条件
	}
	total := first + second // 总和
	length := len(total)
	if total == string(num[:length]) { // 相等则证明目前合法，继续搜索后续
		return dfs(num[length:], second, total)
	}
	// 到最后就是一直不相等，一直小于
	return false
}

// @lc code=end

