/*
 * @lc app=leetcode.cn id=258 lang=golang
 *
 * [258] 各位相加
 */

// @lc code=start
func addDigits(num int) int {
	if num < 10 {
		return num
	}
	for num >= 10 {
		temp := 0
		for num >= 10 {
			temp += num % 10
			num = num / 10
		}
		temp += num
		num = temp
	}
	return num
	// 数学
	// return (num - 1) % 9 + 1
}

// @lc code=end

