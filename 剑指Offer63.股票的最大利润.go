// 剑指 Offer 63. 股票的最大利润
// 假设把某股票的价格按照时间先后顺序存储在数组中，请问买卖该股票一次可能获得的最大利润是多少？

// 示例 1:

// 输入: [7,1,5,3,6,4]
// 输出: 5
// 解释: 在第 2 天（股票价格 = 1）的时候买入，在第 5 天（股票价格 = 6）的时候卖出，最大利润 = 6-1 = 5 。
//      注意利润不能是 7-1 = 6, 因为卖出价格需要大于买入价格。
// 示例 2:

// 输入: [7,6,4,3,1]
// 输出: 0
// 解释: 在这种情况下, 没有交易完成, 所以最大利润为 0。

// 限制：

// 0 <= 数组长度 <= 10^5

// 执行用时：4 ms, 在所有 Go 提交中击败了88.56%的用户
// 内存消耗：2.9 MB, 在所有 Go 提交中击败了100.00%的用户
func maxProfit(prices []int) int {
	// 从右边遍历，找出买后的最高价格
	m := len(prices)
	if m <= 1 {
		return 0
	}
	res := 0
	maxPrice := prices[m-1]
	for i := m - 1; i >= 0; i-- {
		if prices[i] < maxPrice {
			res = max(maxPrice-prices[i], res)
		} else if prices[i] > maxPrice {
			maxPrice = prices[i]
		}
	}
	return res
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}