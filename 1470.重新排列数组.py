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

# 2022-08-29
class Solution:
    def shuffle(self, nums: List[int], n: int) -> List[int]:
        res = [0] * (2 * n)
        for i in range(n):
            res[2 * i] = nums[i]
            res[2 * i + 1] = nums[n + i]
        return res
# @lc code=end