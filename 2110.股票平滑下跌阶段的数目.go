/*
 * @lc app=leetcode.cn id=2110 lang=golang
 *
 * [2110] 股票平滑下跌阶段的数目
 */

// @lc code=start
func getDescentPeriods(prices []int) int64 {
	dp := 1
	var res int64 = 1
	for i := 1; i < len(prices); i++ {
		if prices[i-1] - prices[i] == 1 {
			dp++
		} else {
			dp = 1
		}
		res += int64(dp)
	}
	return res
}
// @lc code=end

