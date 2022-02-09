#
# @lc app=leetcode.cn id=2006 lang=python3
#
# [2006] 差的绝对值为 K 的数对数目
#

# @lc code=start
class Solution:
    def countKDifference(self, nums: List[int], k: int) -> int:
        length = len(nums)
        res = 0
        for i in range(length):
            for j in range (i+1, length):
                if abs(nums[i]-nums[j]) == k:
                    res += 1
        return res
# @lc code=end

