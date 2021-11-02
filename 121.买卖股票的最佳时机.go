/*
 * @lc app=leetcode.cn id=121 lang=golang
 *
 * [121] 买卖股票的最佳时机
 */

// @lc code=start
func maxProfit(prices []int) (profit int) {
	cost := prices[0]
	pricesLength := len(prices)
	for i := 1; i < pricesLength; i++ {
		if prices[i] < cost {
			cost = prices[i] // 比较取得最小买入价格
		}
		if prices[i]-cost > profit {
			profit = prices[i] - cost // 依次比较取得最大收益
		}
	}
	return
}

// @lc code=end

