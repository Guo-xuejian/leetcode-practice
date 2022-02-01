#
# @lc app=leetcode.cn id=1763 lang=python3
#
# [1763] 最长的美好子字符串
#

# @lc code=start
class Solution:
    def longestNiceSubstring(self, s: str) -> str:
        n = len(s)
        maxPos, maxLen = 0, 0
        for i in range(n):
            # lower 和 upper 实际上存储的是当前所有的小写字母和大写字母，只不过分别转化为小写后分在不同位置
            # 如果满足条件,那么 lower == upper
            lower, upper = 0, 0
            for j in range(i, n):
                if s[j].islower():
                    # 做或运算,该字符不做重复计数,出现即可
                    lower |= 1 << (ord(s[j]) - ord('a'))
                else:
                    upper |= 1 << (ord(s[j]) - ord('A'))
                if lower == upper and j - i + 1 > maxLen:   # 每一个起点对应的当前最长美好子字符串
                    maxPos = i
                    maxLen = j - i + 1
        return s[maxPos: maxPos + maxLen]
# @lc code=end
