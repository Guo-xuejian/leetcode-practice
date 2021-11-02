#
# @lc app=leetcode.cn id=66 lang=python3
#
# [66] 加一
#

# @lc code=start
class Solution:
    def plusOne(self, digits: List[int]) -> List[int]:
        digits_length = len(digits)
        for i in range(digits_length-1,-1,-1):
            if digits[i] == 9:  # 只有该位为 9 时才会出现进位
                digits[i] = 0
                if i == 0:      # 首位为 9 则需要改变列表
                    return [1] + digits
            else:
                digits[i] += 1  # 其余情况加一返回
                return digits
        return digits

# @lc code=end

