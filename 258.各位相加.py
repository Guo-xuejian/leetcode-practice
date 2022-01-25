#
# @lc app=leetcode.cn id=258 lang=python3
#
# [258] 各位相加
#

# @lc code=start
class Solution:
    def addDigits(self, num: int) -> int:
        if num < 10:
            return num
        while num >= 10:
            temp = 0
            while num >= 10:
                temp += num % 10    # 加上目前的个位
                num = int(num / 10) # 更新个位
            temp += num
            num = temp
        return num
        # return (num - 1) % 9 + 1
# @lc code=end

