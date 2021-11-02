#
# @lc app=leetcode.cn id=121 lang=python3
#
# [121] 买卖股票的最佳时机
#

# @lc code=start
class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        cost, profit = float("+inf"), 0     # cost 取较大值
        for price in prices:
            cost = min(cost, price)     # 比较取最小买入价格
            profit = max(profit, price - cost)  # 依次比较取得最大收益
        return profit

# @lc code=end

