#
# @lc app=leetcode.cn id=122 lang=python3
#
# [122] 买卖股票的最佳时机 II
#

# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        profit = 0
        for i in range(1, len(prices)):
            temp = prices[i] - prices[i - 1]
            if temp > 0:
                # 只要存在正值收益就卖，同时买入今天的价格(取决于后续是否上涨，不上涨就不买)
                profit += temp
        return profit

# @lc code=end

