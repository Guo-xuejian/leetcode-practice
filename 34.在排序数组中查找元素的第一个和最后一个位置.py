#
# @lc app=leetcode.cn id=34 lang=python3
#
# [34] 在排序数组中查找元素的第一个和最后一个位置
#

# @lc code=start
class Solution:
    def searchRange(self, nums: List[int], target: int) -> List[int]:
        start_index, end_index = -1, -1
        for i in range(len(nums)):
            if nums[i] == target and start_index == -1:
                start_index = i
                end_index = i
            elif nums[i] == target:
                end_index = i
            elif start_index > 0 and end_index > 0 and nums[i] != target:
                break
        return [start_index, end_index]
# @lc code=end

