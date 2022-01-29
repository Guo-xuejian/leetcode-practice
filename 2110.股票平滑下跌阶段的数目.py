#
# @lc app=leetcode.cn id=2110 lang=python3
#
# [2110] 股票平滑下跌阶段的数目
#

# @lc code=start
class Solution:
    def getDescentPeriods(self, prices: List[int]) -> int:
        dp, res = 1, 1  # len(prices) >= 1 所以至少 1

        for i in range(1, len(prices)):
            # dp + 1 意味着与之前满足条件的所有情况两两组合再加上当前情况单独出现 1（else就是只管单独出现并重置数量）
            dp = dp + 1 if prices[i-1] - prices[i] == 1 else 1
            # 已有组合数量加上当前组合数量
            res += dp
        return res
# @lc code=end

