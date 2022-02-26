#
# @lc app=leetcode.cn id=1768 lang=python3
#
# [1768] 交替合并字符串
#

# @lc code=start
class Solution:
    def mergeAlternately(self, word1: str, word2: str) -> str:
        index1, index2 = 0, 0
        m, n = len(word1), len(word2)
        res = ["" for _ in range(m + n)]
        write_index = 0
        while index1 < m or index2 < n:
            if index1 < m:
                res[write_index] = word1[index1]
                index1 += 1
                write_index += 1
            if index2 < n:
                res[write_index] = word2[index2]
                index2 += 1
                write_index += 1
        return "".join(res)
# @lc code=end

