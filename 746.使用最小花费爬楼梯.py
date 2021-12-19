#
# @lc app=leetcode.cn id=746 lang=python3
#
# [746] 使用最小花费爬楼梯
#

# @lc code=start
class Solution:
    def minCostClimbingStairs(self, cost: List[int]) -> int:
        cost_length = len(cost)
        # pre 代表前两位，curr 代表前一个
        pre = curr = 0
        for i in range(2, cost_length + 1):
            # 每一次都取代价最小值，动态规划，对已知条件的遍历，取出最合适的方案
            pre, curr = curr, min(curr + cost[i - 1], pre + cost[i - 2])
        return curr
# @lc code=end

