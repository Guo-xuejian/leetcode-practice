#
# @lc app=leetcode.cn id=13 lang=python3
#
# [13] 罗马数字转整数
#

# @lc code=start
class Solution:
    def romanToInt(self, s: str) -> int:
        if not s:
            return 0
        symbols_values_dict = {
            "M": 1000,
            "D": 500,
            "C": 100,
            "L": 50,
            "X": 10,
            "V": 5,
            "I": 1,
        }
        s_length = len(s)
        # 初始罗马数值 num, 临时数值记录上一个数字 lastNumber,返回值 total
        num = last_number = total = 0
        # 倒序遍历,从最小值开始往上加,遇到组合的那种也很好处理,减掉较小值就完成了实际的加总,比如 CM 900,从后往前先加上 1000,遇到 C 之后因为它比 M 小,total - 100 即可
        for i in range(len(s)-1,-1,-1):
            char = s[i]
            num = symbols_values_dict[char]
            # 遇到了组合数字,需要在 toatal 的基础上减去这个值
            # 同时题目中给出的条件只会出现一个数字为两个字符的组合
            # 那么需要一个 lastNumber 记录
            if num < last_number:
                total -= num
            else:
                total += num
            # 记录这一次遍历的值,实际上不需要担心 I,V,X等对M的影响,量级有差异,不会出现
            last_number = num
        return total

# @lc code=end
