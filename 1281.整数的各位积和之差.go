/*
 * @lc app=leetcode.cn id=1281 lang=golang
 *
 * [1281] 整数的各位积和之差
 */

// @lc code=start
func subtractProductAndSum(n int) int {
	multiNum, sumNum := 1, 0
	for n > 0 {
		curr := n % 10
		multiNum *= curr
		sumNum += curr
		n = n / 10
	}
	return multiNum - sumNum
}
// @lc code=end

