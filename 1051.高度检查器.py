#
# @lc app=leetcode.cn id=1051 lang=python3
#
# [1051] 高度检查器
#

# @lc code=start
class Solution:
    def heightChecker(self, heights: List[int]) -> int:
        # expected = sorted(heights)
        # return sum(1 for x, y in zip(heights, expected) if x != y)
        expected = sorted(heights)
        res = 0
        for i in range(len(heights)):
            if expected[i] != heights[i]:
                res += 1
        return res
# @lc code=end

