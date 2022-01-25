#
# @lc app=leetcode.cn id=561 lang=python3
#
# [561] 数组拆分 I
#

# @lc code=start
class Solution:
    def arrayPairSum(self, nums: List[int]) -> int:
        res = 0
        for i in range(0, len(nums), 2):
            res += nums[i]
        return res
# @lc code=end

