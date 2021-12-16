#
# @lc app=leetcode.cn id=1470 lang=python3
#
# [1470] 重新排列数组
#

# @lc code=start
class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        result = []
        for i in range(n):
            result.append(nums[i])
            result.append(nums[i+n])
        return result
# @lc code=end