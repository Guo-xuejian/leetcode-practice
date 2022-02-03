#
# @lc app=leetcode.cn id=504 lang=python3
#
# [504] 七进制数
#

# @lc code=start
class Solution:
    def convertToBase7(self, num: int) -> str:
        base = 7    # 进制
        if not num: # 为 0
            return str(num)

        sign = num < 0  # 正负判断
        num = abs(num)  # 去正
        res = ""
        while num:
            # 在 res 前加入新字符
            res = str(num % base) + res
            # 辗转相除
            num //= base
        # 加上符号
        if sign:
            res = "-" + res
        return res
# @lc code=end

