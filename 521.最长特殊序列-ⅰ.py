#
# @lc app=leetcode.cn id=521 lang=python3
#
# [521] 最长特殊序列 Ⅰ
#

# @lc code=start
class Solution:
    def findLUSlength(self, a: str, b: str) -> int:
        # 其实只有相等和不相等两种
        # 相等则 -1，不相等则取较长长度，因为较长的那个字符串必定不是另一个的子序列
        if a == b:
            return -1
        return max(len(a), len(b))
# @lc code=end

