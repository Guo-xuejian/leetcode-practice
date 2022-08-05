#
# @lc app=leetcode.cn id=1408 lang=python3
#
# [1408] 数组中的字符串匹配
#

# @lc code=start
class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:
        length, res = len(words), []

        for i in range(length):
            for j in range(length):
                if i != j and words[i] in words[j]:
                    res.append(words[i])
                    break
        return res

# 2022-08-06
class Solution:
    def stringMatching(self, words: List[str]) -> List[str]:
        exist = {}
        res = []
        for i, word in enumerate(words):
            for j, item in enumerate(words):
                if i == j or exist.__contains__(item):
                    continue
                if item in word:
                    res.append(item)
                    exist[item] = ""
        return res
# @lc code=end
