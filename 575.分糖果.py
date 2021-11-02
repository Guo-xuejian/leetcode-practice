#
# @lc app=leetcode.cn id=575 lang=python3
#
# [575] 分糖果
#

# @lc code=start
class Solution:
    def distributeCandies(self, candyType: List[int]) -> int:
        # 去重比较
        return min(len(set(candyType)), int(len(candyType)/2))
# @lc code=end


