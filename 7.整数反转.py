#
# @lc app=leetcode.cn id=7 lang=python3
#
# [7] 整数反转
#

# @lc code=start
class Solution:
    def reverse(self, x: int) -> int:
        # 首先确定数字是否为负数，因为负数需要处理符号
        if x == 0:
            return 0
        if x > 0:
            # 正数
            res = int(str(x)[::-1])
            if res > 2**31 - 1:
                return 0
            return res
        else:
            # 负数
            res = int(str(x)[1:][::-1])
            if res > 2**31:
                return 0
            else:
                return res*(-1)
# @lc code=end

