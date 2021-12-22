#
# @lc app=leetcode.cn id=1154 lang=python3
#
# [1154] 一年中的第几天
#

# @lc code=start
class Solution:
    def dayOfYear(self, date: str) -> int:
        # 拆出年月日
        year, month, day = [int(x) for x in date.split("-")]

        # 模拟每一年的每月天数
        amount = [31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        # 闰年二月加一天
        if year % 400 == 0 or (year % 4 == 0 and year % 100 != 0):
            amount[1] += 1
        # 加总前几个月并加上本月的天数
        return sum(amount[:month - 1]) + day
# @lc code=end

