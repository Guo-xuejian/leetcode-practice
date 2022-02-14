#
# @lc app=leetcode.cn id=1754 lang=python3
#
# [1754] 构造字典序最大的合并字符串
#

# @lc code=start
class Solution:
    def largestMerge(self, word1: str, word2: str) -> str:
        # 注释参考 go 代码
        def greater_dictionary_order(a, b):
            m, n = len(a), len(b)
            for i in range(m):
                if i + 1 > n or a[i] > b[i]:
                    return True
                elif a[i] < b[i]:
                    return False
            return False

        m, n = len(word1), len(word2)
        all_letter = ["" for _ in range(m + n)]
        idx, idx1, idx2 = 0, 0, 0
        while idx1 < m and idx2 < n:
            if greater_dictionary_order(word1[idx1:], word2[idx2:]):
                all_letter[idx] = word1[idx1]
                idx1 += 1
            else:
                all_letter[idx] = word2[idx2]
                idx2 += 1
            idx += 1

        while idx1 < m:
            all_letter[idx] = word1[idx1]
            idx1 += 1
            idx += 1
        while idx2 < n:
            all_letter[idx] = word2[idx2]
            idx2 += 1
            idx += 1
        return "".join(all_letter)

# @lc code=end
