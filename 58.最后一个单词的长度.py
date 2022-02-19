#
# @lc app=leetcode.cn id=58 lang=python3
#
# [58] 最后一个单词的长度
#

# @lc code=start
class Solution:
    def lengthOfLastWord(self, s: str) -> int:
        res = 0
        for i in range(len(s) - 1, -1, -1):
            if s[i] != ' ':
                res += 1
            else:
                if res > 0:
                    break
        return res
# @lc code=end
