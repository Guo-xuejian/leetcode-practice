#
# @lc app=leetcode.cn id=1189 lang=python3
#
# [1189] “气球” 的最大数量
#

# @lc code=start
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        letter_dict = {"b": 0, "a": 0, "l": 0, "o": 0, "n": 0}
        for letter in text:
            if letter in letter_dict:
                letter_dict[letter] += 1
        res = len(text)
        for key, val in letter_dict.items():
            if key == "o" or  key == "l":
                res = min(res, val // 2)
            else:
                res = min(res, val)
        return res
# @lc code=end
