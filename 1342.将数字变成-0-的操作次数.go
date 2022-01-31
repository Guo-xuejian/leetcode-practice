/*
 * @lc app=leetcode.cn id=1342 lang=golang
 *
 * [1342] 将数字变成 0 的操作次数
 */

// @lc code=start
func numberOfSteps(num int) (res int) {
	for num != 0 {
		if num % 2 != 0 {
			num--
		} else {
			num /= 2
		}
		res++
	}
	return
}
// @lc code=end

