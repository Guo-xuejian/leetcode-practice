#
# @lc app=leetcode.cn id=389 lang=python3
#
# [389] 找不同
#

# @lc code=start
class Solution:
    def findTheDifference(self, s: str, t: str) -> str:
        length = len(s)
        if length == 0:
            return t[0]
        letter_dict = {}
        index = 0
        while index < length:
            if s[index] in letter_dict:
                letter_dict[s[index]] += 1
            else:
                letter_dict[s[index]] = 1

            if t[index] in letter_dict:
                letter_dict[t[index]] -= 1
            else:
                letter_dict[t[index]] = -1
            index += 1
        if t[length] not in letter_dict:
            return t[length]
        else:
            letter_dict[t[length]] -= 1
        for key, val in letter_dict.items():
            if val == -1:
                return key
# @lc code=end
