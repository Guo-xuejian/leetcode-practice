#
# @lc app=leetcode.cn id=1608 lang=python3
#
# [1608] 特殊数组的特征值
#

# @lc code=start
class Solution:
    def specialArray(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(1, len(nums) + 1):
            if nums[-i] < i:
                if nums[-i] < i - 1:
                    return (i - 1)
                else:
                    return -1
        return len(nums)
# @lc code=end

