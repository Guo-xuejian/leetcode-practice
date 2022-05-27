#
# @lc app=leetcode.cn id=1021 lang=python3
#
# [1021] 删除最外层的括号
#

# @lc code=start
class Solution:
    def removeOuterParentheses(self, s: str) -> str:
        res, stack = "", []
        for c in s:
            if c == ')':
                stack.pop()
            if stack:
                res += c
            if c == '(':
                stack.append(c)
        return res
# @lc code=end

