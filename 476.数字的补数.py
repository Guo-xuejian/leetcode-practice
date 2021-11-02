#
# @lc app=leetcode.cn id=476 lang=python3
#
# [476] 数字的补数
#

# @lc code=start
class Solution:
    def findComplement(self, num: int) -> int:
        highbit = 0
        for i in range(1, 31):
            if num >= (1 << i):
                highbit += 1
            else:
                break
        mask = (1 << (highbit + 1)) - 1
        return num ^ mask
# @lc code=end

