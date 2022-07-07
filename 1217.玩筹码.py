#
# @lc app=leetcode.cn id=1217 lang=python3
#
# [1217] 玩筹码
#

# @lc code=start
class Solution:
    def minCostToMoveChips(self, position: List[int]) -> int:
        # 偶数位是等价的，奇数位也可以通过移动偶数位等价
        # 所以实际上偶数位上的点可以看做同一个点，奇数位同样，比较奇偶数那个小一些，移动的代价就最小
        odd_num = sum(pos % 2 for pos in position)
        return min(odd_num, len(position) - odd_num)


# class Solution:
#     def minCostToMoveChips(self, position: List[int]) -> int:
#         cnt = Counter(p % 2 for p in position)  # 根据模 2 后的余数来统计奇偶个数
#         return min(cnt[0], cnt[1])
# @lc code=end

