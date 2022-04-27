#
# @lc app=leetcode.cn id=905 lang=python3
#
# [905] 按奇偶排序数组
#

# @lc code=start
class Solution:
    def sortArrayByParity(self, nums: List[int]) -> List[int]:
        nums.sort(key = lambda e: e % 2 != 0)
        return nums
# @lc code=end

