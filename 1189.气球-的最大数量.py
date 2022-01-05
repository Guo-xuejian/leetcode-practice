#
# @lc app=leetcode.cn id=1189 lang=python3
#
# [1189] “气球” 的最大数量
#

# @lc code=start
class Solution:
    def maxNumberOfBalloons(self, text: str) -> int:
        letter_list = ["b", "a", "l", "o", "n"]
        letter_dict = {}
        for letter in letter_list:  # 通过固定字符判定是否结束，同时初始化一个字典用来记录 balloon 中的字符个数
            if letter not in text:
                return 0
            else:
                letter_dict[letter] = 0
        for one in text:    # 记录 balloon 中字符出现次数
            if one not in letter_list:
                continue
            letter_dict[one] += 1
        # 双数字符为 l o
        double_limit = min(
            [val for key, val in letter_dict.items() if key in ["l", "o"]])
        if double_limit < 2:
            return 0
        # 单数字符为 b a n
        single_limit = min(
            [val for key, val in letter_dict.items() if key in ["a", "b", "n"]])

        return min(int(double_limit/2), single_limit)   # 全部转化为单数后比较最小值
# @lc code=end
