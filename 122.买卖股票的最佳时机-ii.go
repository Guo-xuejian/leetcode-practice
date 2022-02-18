/*
 * @lc app=leetcode.cn id=122 lang=golang
 *
 * [122] 买卖股票的最佳时机 II
 */

// @lc code=start
func maxProfit(prices []int) (profit int) {
	for i := 1; i < len(prices); i++ {
		temp := prices[i] - prices[i - 1]
		if temp > 0 {
			// 只要存在正值收益就卖，同时买入今天的价格(取决于后续是否上涨，不上涨就不买)
			profit += temp
		}
	}
	return
}
// @lc code=end

