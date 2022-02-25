#
# @lc app=leetcode.cn id=2016 lang=python3
#
# [2016] 增量元素之间的最大差值
#

# @lc code=start
class Solution:
    def maximumDifference(self, nums: List[int]) -> int:
        # n = len(nums)
        # res = -1
        # for i in range(n):
        #     for j in range(i + 1, n):
        #         if nums[j] > nums[i] and nums[j] - nums[i] > res:
        #             res = nums[j] - nums[i]
        # return res
        n = len(nums)
        res, premin = -1, nums[0]

        for i in range(1, n):
            if nums[i] > premin:
                res = max(res, nums[i] - premin)
            else:
                premin = nums[i]

        return res
# @lc code=end
