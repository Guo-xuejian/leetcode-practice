#
# @lc app=leetcode.cn id=43 lang=python3
#
# [43] 字符串相乘
#

# @lc code=start
class Solution:
    def multiply(self, num1: str, num2: str) -> str:
        if num1 == "0" or num2 == "0":
            return "0"
        return str(self.getInteger(num1) * self.getInteger(num2))
        
    def getInteger(self, s: str) -> int:
        res_num = 0
        # s_length = max_bit = len(s)
        s_length = len(s)
        for i in range(s_length):   # 根据进位计算得出该数字
            res_num += int(s[i]) * 10**(s_length-i-1)
            # max_bit -= 1
        return res_num
# @lc code=end

