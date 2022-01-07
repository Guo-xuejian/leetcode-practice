#
# @lc app=leetcode.cn id=1614 lang=python3
#
# [1614] 括号的最大嵌套深度
#

# @lc code=start
class Solution:
    def maxDepth(self, s: str) -> int:
        res = 0
        left = 0    # 不使用栈，节省点空间
        # 实际上也只需要栈的长度这个属性
        for one in s:
            if one == "(":  # 左括号数量 +1 后比较嵌套深度
                left += 1
                if left > res:
                    res = left
            elif one == ")":    # 右括号抵消一个左括号，由于题目给的一定是合法嵌套，所以不担心，但还是写一下吧
                left = left -1 if left > 0 else 0
        return res
# @lc code=end

