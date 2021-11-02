#
# @lc app=leetcode.cn id=282 lang=python3
#
# [282] 给表达式添加运算符
#

# @lc code=start
class Solution:
    def addOperators(self, num: str, target: int) -> List[str]:
        n = len(num)
        ans = []

        def backtrack(expr, i, res, nul):
            if i == n:
                if res == target:
                    ans.append("".join(expr))
                return
            sign_index = len(expr)
            if i > 0:
                expr.append("")     # 占位，下面填充符号
            val = 0
            for j in range(i, n):    # 枚举截取的数字长度（取多少位）
                if j > i and num[i] == "0":     # 数字可以是单个0，但不能有前导的0
                    break
                # 不一定都是一位数字，多位数字的算式组合也是符合要求的
                val = val*10 + int(num[j])      # 数字进位处理
                expr.append(num[j])     # 数字进入
                if i == 0:      #表达式开头不能添加符号
                    backtrack(expr, j+1, val, val)
                else:           # 枚举符号
                    expr[sign_index] = "+"
                    backtrack(expr, j+1, res+val, val)
                    expr[sign_index] = "-"
                    backtrack(expr, j+1, res-val, -val)
                    expr[sign_index] = "*"
                    # 若添加 * 号，由于乘法运算优先级高于加法和减法运算,需要撤销上一次的运算结果，直接减去，然后加上新的乘法结果
                    backtrack(expr, j+1, res-nul+nul*val, nul*val)
                # 一次遍历结束去除刚才的算式内容，因为一直再写入，但是需要expr 保持为空，所以每次最后都需要删除全部元素
            del expr[sign_index:]
        backtrack([], 0, 0, 0)
        return ans
# @lc code=end

