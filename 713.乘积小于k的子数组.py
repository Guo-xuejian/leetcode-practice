#
# @lc app=leetcode.cn id=713 lang=python3
#
# [713] 乘积小于K的子数组
#

# @lc code=start
class Solution:
    def numSubarrayProductLessThanK(self, nums: List[int], k: int) -> int:
        ans, prod, i = 0, 1, 0
        for j, num in enumerate(nums):
            prod *= num
            while i <= j and prod >= k:
                prod //= nums[i]
                i += 1
            ans += j - i + 1
        return ans
# @lc code=end

