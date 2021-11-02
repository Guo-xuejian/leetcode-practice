#
# @lc app=leetcode.cn id=441 lang=python3
#
# [441] 排列硬币
#

# @lc code=start
class Solution:
    def arrangeCoins(self, n: int) -> int:
        if n <= 1:
            return n
        totalCoins = n
        for i in range(n):
            totalCoins -= i + 1
            if totalCoins >= 0:
                continue
            else:
                return i
# @lc code=end

