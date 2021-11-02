/*
 * @lc app=leetcode.cn id=171 lang=golang
 *
 * [171] Excel 表列序号
 */

// @lc code=start
func titleToNumber(columnTitle string) int {
	num, multiple := 0, 1
	length := len(columnTitle)
	for i := length - 1; i >= 0; i-- { // 反向 从低位开始出发
		k := int(columnTitle[i]) - 64 // 遍历即是 rune 型，直接转化为对应的 ascill 码值获取题目对应的序号
		num += k * multiple           // 对应倍数
		multiple *= 26                // 倍数随位数递增
	}
	return num
}

// @lc code=end

