#
# @lc app=leetcode.cn id=53 lang=python3
#
# [53] 最大子数组和
#

# @lc code=start
class Solution:
    def maxSubArray(self, nums: List[int]) -> int:
        length = len(nums)
        pre = 0
        res = nums[0]
        for i in range(length):
            pre = max(nums[i], nums[i] + pre)
            res = max(pre, res)
        return res
# @lc code=end

