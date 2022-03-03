#
# @lc app=leetcode.cn id=2104 lang=python3
#
# [2104] 子数组范围和
#

# @lc code=start
class Solution:
    def subArrayRanges(self, nums: List[int]) -> int:
        res, n = 0, len(nums)
        for i in range(n):
            minVal, maxVal = inf, -inf
            for j in range(i, n):
                minVal = min(minVal, nums[j])
                maxVal = max(maxVal, nums[j])
                res += maxVal - minVal
        return res
# @lc code=end

