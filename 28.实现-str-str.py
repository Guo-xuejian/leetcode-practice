#
# @lc app=leetcode.cn id=28 lang=python3
#
# [28] 实现 strStr()
#

# @lc code=start
class Solution:
    def strStr(self, haystack: str, needle: str) -> int:
        length1, length2 = len(haystack), len(needle)

        for i in range(length1 - length2 + 1):
            if haystack[i:i+length2] == needle:
                return i
        
        return -1
# @lc code=end

