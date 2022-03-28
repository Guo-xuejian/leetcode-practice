#
# @lc app=leetcode.cn id=693 lang=python3
#
# [693] 交替位二进制数
#

# @lc code=start
class Solution:
    def hasAlternatingBits(self, n: int) -> bool:
        a = n ^ (n >> 1)
        return a & (a + 1) == 0
# @lc code=end

