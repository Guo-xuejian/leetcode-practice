#
# @lc app=leetcode.cn id=1342 lang=python3
#
# [1342] 将数字变成 0 的操作次数
#

# @lc code=start
class Solution:
    def numberOfSteps(self, num: int) -> int:
        res = 0
        while num != 0:
            if num % 2 != 0:
                num -= 1
            else:
                num = int(num / 2)
            res += 1
        return res
# @lc code=end

