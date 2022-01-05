#
# @lc app=leetcode.cn id=2057 lang=python3
#
# [2057] 值相等的最小索引
#

# @lc code=start
class Solution:
    def smallestEqual(self, nums: List[int]) -> int:
        for i in range(len(nums)):
            if i % 10 == nums[i]:
                return i
        return -1
# @lc code=end

