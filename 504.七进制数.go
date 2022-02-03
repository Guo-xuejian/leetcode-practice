/*
 * @lc app=leetcode.cn id=504 lang=golang
 *
 * [504] 七进制数
 */

// @lc code=start
func convertToBase7(num int) (res string) {
	base := 7
	if num == 0 {
		res = "0"
		return
	}

	sign := num < 0 // 正负判断
	num = abs(num)  // 计算使用正值
	for num != 0 {
		res = strconv.Itoa(num%base) + res // 在前面写入对应字符
		num /= base                        // 辗转相除
	}
	if sign { // 为负数加上符号
		res = "-" + res
	}
	return
}

func abs(num int) int {
	if num < 0 {
		return -num
	}
	return num
}

// @lc code=end

