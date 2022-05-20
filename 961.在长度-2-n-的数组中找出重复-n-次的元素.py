#
# @lc app=leetcode.cn id=961 lang=python3
#
# [961] 在长度 2N 的数组中找出重复 N 次的元素
#

# @lc code=start
class Solution:
    def repeatedNTimes(self, nums: List[int]) -> int:
        nums.sort()
        if nums[-1] == nums[len(nums) // 2]:
            return nums[-1]
        else:
            return nums[len(nums) // 2 - 1]
# @lc code=end

