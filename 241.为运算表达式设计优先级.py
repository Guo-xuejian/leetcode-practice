#
# @lc app=leetcode.cn id=241 lang=python3
#
# [241] 为运算表达式设计优先级
#

# @lc code=start
ADDITION = -1
SUBTRACTION = -2
MULTIPLICATION = -3

class Solution:
    def diffWaysToCompute(self, expression: str) -> List[int]:
        ops = []
        i, n = 0, len(expression)
        while i < n:
            if expression[i].isdigit():
                x = 0
                while i < n and expression[i].isdigit():
                    x = x * 10 + int(expression[i])
                    i += 1
                ops.append(x)
            else:
                if expression[i] == '+':
                    ops.append(ADDITION)
                elif expression[i] == '-':
                    ops.append(SUBTRACTION)
                else:
                    ops.append(MULTIPLICATION)
                i += 1

        @cache
        def dfs(l: int, r: int) -> List[int]:
            if l == r:
                return [ops[l]]
            res = []
            for i in range(l, r, 2):
                left = dfs(l, i)
                right = dfs(i + 2, r)
                for x in left:
                    for y in right:
                        if ops[i + 1] == ADDITION:
                            res.append(x + y)
                        elif ops[i + 1] == SUBTRACTION:
                            res.append(x - y)
                        else:
                            res.append(x * y)
            return res
        return dfs(0, len(ops) - 1)
# @lc code=end

