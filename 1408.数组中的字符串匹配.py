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
# @lc code=end
