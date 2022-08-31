/*
 * @lc app=leetcode.cn id=1475 lang=golang
 *
 * [1475] 商品折扣后的最终价格
 */

// @lc code=start
func finalPrices(prices []int) []int {
	ans := make([]int, len(prices))
	for i, p := range prices {
		discount := 0
		for _, q := range prices[i+1:] {
			if q <= p {
				discount = q
				break
			}
		}
		ans[i] = p - discount
	}
	return ans
}

// @lc code=end

