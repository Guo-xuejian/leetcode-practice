#
# @lc app=leetcode.cn id=1185 lang=python3
#
# [1185] 一周中的第几天
#

# @lc code=start
class Solution:
    def dayOfTheWeek(self, day: int, month: int, year: int) -> str:
        #----1971年1月1日是周五

        month_day = [0, 31, 28, 31, 30, 31, 30, 31, 31, 30, 31, 30, 31]
        day_week = ["Sunday", "Monday", "Tuesday", "Wednesday", "Thursday", "Friday", "Saturday"]

        cur = 0
        for y in range(1971, year):
            cur += 365
            if self.is_leap_year(y) == True:
                cur += 1
        for m in range(month):
            cur += month_day[m]
            if m == 2 and self.is_leap_year(year) == True:
                cur += 1
        cur += day

        res = day_week[(cur + 4) % 7]
        return res


    def is_leap_year(self, year: int) -> bool:
        return (year % 4 == 0 and year % 100 != 0) or (year % 400 == 0)
# @lc code=end

