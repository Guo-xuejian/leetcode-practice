#
# @lc app=leetcode.cn id=152 lang=python3
#
# [152] 乘积最大子数组
#

# @lc code=start
class Solution:
    def maxProduct(self, nums: List[int]) -> int:
        max_num, min_num, res = nums[0], nums[0], nums[0]
        for i in range(1, len(nums)):
            ma, mi = max_num, min_num
            max_num = max(ma * nums[i], max(nums[i], mi * nums[i]))
            min_num = min(mi * nums[i], min(nums[i], ma * nums[i]))
            res = max(res, max_num)
        return res

# @lc code=end

