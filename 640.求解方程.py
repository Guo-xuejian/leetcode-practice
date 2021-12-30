#
# @lc app=leetcode.cn id=640 lang=python3
#
# [640] 求解方程
#

# @lc code=start
class Solution:
    def solveEquation(self, equation: str) -> str:
        def split_line(line):
            if line[0] == "x":
                line = "1" + line   # x 为第一个字符需要加上个 1
            # 全部替换为以 + 号组织的字符串，最后以 + 切分
            text = line.replace("+x", "+1x").replace("-x",
                                                     "-1x").replace("-", "+-").split("+")
            #
            text = [one for one in text if len(one) > 0]

            x = sum([int(t[:-1]) for t in text if t[-1] == "x"])
            num = sum([int(t) for t in text if t[-1] != "x"])
            return x, num
        line_left, line_right = equation.split("=")
        x_left, num_left = split_line(line_left)
        x_right, num_right = split_line(line_right)
        x = x_left - x_right    # 最终变量系数
        num = num_right - num_left  # 最终常数
        if x == 0:
            if num == 0:
                return "Infinite solutions"
            else:
                return "No solution"
        else:
            return "x=" + str(num // x)  # 要求整数

# @lc code=end
