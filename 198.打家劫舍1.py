#
# @lc app=leetcode.cn id=198 lang=python3
#
# [198] 打家劫舍
#

# @lc code=start
class Solution:
    def rob(self, nums: List[int]) -> int:
        length = len(nums)
        if length == 1:
            return nums[0]
        pre_num, curr_num = nums[0], max(nums[0], nums[1])
        for i in range(2, length):
            pre_num, curr_num = curr_num, max(pre_num + nums[i], curr_num)
        return curr_num
# @lc code=end
