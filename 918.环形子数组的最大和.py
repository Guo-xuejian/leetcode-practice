#
# @lc app=leetcode.cn id=918 lang=python3
#
# [918] 环形子数组的最大和
#

# @lc code=start
class Solution:
    def maxSubarraySumCircular(self, nums: List[int]) -> int:
        total, max_sum, curr_max, min_sum, curr_min = 0, nums[0], 0, nums[0], 0

        for num in nums:
            total += num
            curr_max = max(curr_max + num, num)
            max_sum = max(max_sum, curr_max)

            curr_min = min(curr_min + num, num)
            min_sum = min(min_sum, curr_min)
        return max(max_sum, total - min_sum) if max_sum > 0 else max_sum
# @lc code=end

