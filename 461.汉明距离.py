#
# @lc app=leetcode.cn id=461 lang=python3
#
# [461] 汉明距离
#

# @lc code=start
class Solution:
    def hammingDistance(self, x: int, y: int) -> int:
        temp, res = x ^ y, 0
        while temp > 0:
            res += temp & 1
            temp >>= 1
        return res
# @lc code=end

