#
# @lc app=leetcode.cn id=462 lang=python3
#
# [462] 最少移动次数使数组元素相等 II
#

# @lc code=start
class Solution:
    def minMoves2(self, nums: List[int]) -> int:
        nums.sort()
        middle_num = nums[len(nums) // 2]
        return sum(abs(num - middle_num) for num in nums)
# @lc code=end

