#
# @lc app=leetcode.cn id=1708 lang=python3
#
# [1708] 长度为K的最大子数组

#

# @lc code=start
class Solution:
    def largestSubarray(self, nums: List[int], k: int) -> List[int]:
        border = len(nums) - k + 1
        max_num_index = nums[:border].index(max(nums[:border]))
        return nums[max_num_index:max_num_index + k]

# @lc code=end

