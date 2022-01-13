#
# @lc app=leetcode.cn id=747 lang=python3
#
# [747] 至少是其他数字两倍的最大数
#

# @lc code=start
class Solution:
    def dominantIndex(self, nums: List[int]) -> int:
        # 维护最大值，次大值和最大值索引
        max_num, pre_max_num, index = nums[0], -1, 0
        for i in range(1, len(nums)):
            if nums[i] > max_num:
                max_num, pre_max_num, index = nums[i], max_num, i
            elif nums[i] < max_num and nums[i] > pre_max_num:
                pre_max_num = nums[i]
        # 至少为两倍
        return index if max_num >= pre_max_num * 2 else -1
# @lc code=end
