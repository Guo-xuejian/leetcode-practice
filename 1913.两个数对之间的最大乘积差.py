#
# @lc app=leetcode.cn id=1913 lang=python3
#
# [1913] 两个数对之间的最大乘积差
#

# @lc code=start
class Solution:
    def maxProductDifference(self, nums: List[int]) -> int:
        # 找出最大值和次大值，最小值和次小值
        max_num_index, pre_max_num_index = 0, 0
        min_num_index, pre_min_num_index = 0, 0
        for i in range(1, len(nums)):
            curr = nums[i]
            # 最小值和次小值
            if curr <= nums[min_num_index]:  # 等于也可，有可能出现相同的两个最大值
                min_num_index, pre_min_num_index = i, min_num_index
            elif curr > nums[min_num_index] and (curr < nums[pre_min_num_index] or (pre_min_num_index == 0 and min_num_index == 0)):
                pre_min_num_index = i

            # 最大值和次大值
            if curr >= nums[max_num_index]:  # 同上
                max_num_index, pre_max_num_index = i, max_num_index
            elif curr < nums[max_num_index] and (curr > nums[pre_max_num_index] or (pre_max_num_index == 0 and max_num_index == 0)):
                pre_max_num_index = i
        return nums[max_num_index]*nums[pre_max_num_index] - nums[min_num_index]*nums[pre_min_num_index]
# @lc code=end
