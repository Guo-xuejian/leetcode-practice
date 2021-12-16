#
# @lc app=leetcode.cn id=1014 lang=python3
#
# [1014] 最佳观光组合
#

# @lc code=start
class Solution:
    def maxScoreSightseeingPair(self, values: List[int]) -> int:
        max_score, left_best_sight = 0, values[0] + 0
        for i in range(1, len(values)):
            max_score = max(max_score, left_best_sight+values[i] - i)
            # 一边遍历一边维护前方的景点，取前方景点中值与索引之和最大的
            left_best_sight = max(left_best_sight, values[i] + i)
        return max_score
# @lc code=end

