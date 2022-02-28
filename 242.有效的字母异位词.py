#
# @lc app=leetcode.cn id=242 lang=python3
#
# [242] 有效的字母异位词
#

# @lc code=start
class Solution:
    def isAnagram(self, s: str, t: str) -> bool:
        letter_dict = {}
        index = 0
        while index < min(len(s), len(t)):
            if s[index] not in letter_dict:
                letter_dict[s[index]] = 1
            else:
                letter_dict[s[index]] += 1
            
            if t[index] not in letter_dict:
                letter_dict[t[index]] = -1
            else:
                letter_dict[t[index]] -= 1
            index += 1
        while index < len(s):
            if s[index] not in letter_dict:
                letter_dict[s[index]] = 1
            else:
                letter_dict[s[index]] += 1
            index += 1
        while index < len(t):
            if t[index] not in letter_dict:
                letter_dict[t[index]] = -1
            else:
                letter_dict[t[index]] -= 1
            index += 1
        for val in letter_dict.values():
            if val != 0:
                return False
        return True
# @lc code=end

