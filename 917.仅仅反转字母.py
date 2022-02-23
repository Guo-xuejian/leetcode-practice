#
# @lc app=leetcode.cn id=917 lang=python3
#
# [917] 仅仅反转字母
#

# @lc code=start
class Solution:
    def reverseOnlyLetters(self, s: str) -> str:
        length = len(s)
        if length == 1:
            return s
        temp = list(s)
        left, right = 0, length - 1
        while left < right:
            if temp[left].isalpha():
                while left < right and not temp[right].isalpha():
                    right -= 1
                if right <= left:
                    break
                temp[left], temp[right] = temp[right], temp[left]
                right -= 1
            left += 1
        return "".join(temp)
# @lc code=end
