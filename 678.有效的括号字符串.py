#
# @lc app=leetcode.cn id=678 lang=python3
#
# [678] 有效的括号字符串
#

# @lc code=start
class Solution:
    def checkValidString(self, s: str) -> bool:
        left, aster = [], []
        length = len(s)
        for i in range(length):
            if s[i] == "(":
                left.append(i)
            elif s[i] == "*":
                aster.append(i)
            elif len(left) > 0: # 右括号匹配一个左括号
                left = left[:len(left)-1]
            elif len(aster) > 0:    # 无左括号则匹配一个 *
                aster = aster[:len(aster)-1]
            else:   # 都无则 False
                return False
        i = len(left) - 1
        j = len(aster) - 1
        while i >= 0 and j >= 0:
            if left[i] > aster[j]:
                return False
            i, j = i - 1, j - 1
        # 小于 0 则完整匹配结束
        return i < 0
# @lc code=end

