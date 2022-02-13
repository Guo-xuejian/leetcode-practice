#
# @lc app=leetcode.cn id=213 lang=python3
#
# [213] 打家劫舍 II
#

# @lc code=start
class Solution:
    def rob(self, nums: List[int]) -> int:
        def robRange(start, end):
            if end - start == 1:
                return max(nums[start], nums[end])
            pre_num, curr_num = nums[start], max(nums[start], nums[start + 1])
            for i in range(start + 2, end + 1):
                pre_num, curr_num = curr_num, max(pre_num + nums[i], curr_num)
            return curr_num
        
        length = len(nums)
        if length == 1:
            return nums[0]
        elif length == 2:
            return max(nums[0], nums[1])
        else:
            return max(robRange(0, length - 2), robRange(1, length - 1))
# @lc code=end
