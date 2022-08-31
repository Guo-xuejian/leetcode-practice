#
# @lc app=leetcode.cn id=1475 lang=python3
#
# [1475] 商品折扣后的最终价格
#

# @lc code=start
class Solution:
    def finalPrices(self, prices: List[int]) -> List[int]:
        n = len(prices)
        ans = [0] * n
        for i, p in enumerate(prices):
            discount = 0
            for j in range(i + 1, n):
                if prices[j] <= p:
                    discount = prices[j]
                    break
            ans[i] = p - discount
        return ans
# @lc code=end

