#
# @lc app=leetcode.cn id=12 lang=python3
#
# [12] 整数转罗马数字
#

# @lc code=start
class Solution:
    def intToRoman(self, num: int) -> str:
        # 由于输入值有限制,那么久枚举出所有的可能
        values = [1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1]
        symbols = ["M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"]
        result = ""
        index = 0
        while num != 0:
            while values[index] > num:
                index += 1
            num -= values[index]
            result += symbols[index]
        return result
# @lc code=end
