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


# 2022-08-10
class Solution:
    def solveEquation(self, equation: str) -> str:
        def getNum(string_num):
            if string_num == "-":
                return -1
            elif string_num == "":
                return 1
            res = 0
            multi = 1
            for i in range(len(string_num)-1, -1, -1):
                if string_num[i] == "+":
                    continue
                elif string_num[i] == "-":
                    res *= -1
                else:
                    res += (ord(string_num[i]) - ord("0")) * multi
                    multi *= 10
            return res

        # 合并变量和常量分别在左边和右边
        equation = equation.replace("-", "+-")  # 做替换，方便切分
        left_sting, right_string = equation.split("=")
        var_num, const_num = 0, 0
        # 先处理左边
        left, right = left_sting.split("+"), right_string.split("+")
        curr_length = 0
        for one in left:
            if one == "":
                continue
            curr_length = len(one)
            if one[curr_length - 1] == "x":  # 左变量系数相加
                var_num += getNum(one[:len(one) - 1])
            else:   # 左常量相减
                const_num -= getNum(one)
        
        for one in right:
            if one == "":
                continue
            curr_length = len(one)
            if one[curr_length - 1] == "x":  # 右变量系数相减
                var_num -= getNum(one[:len(one) - 1])
            else:   # 右常量相加
                const_num += getNum(one)
        
        # 都为 0 则无数解
        if var_num == 0 and const_num == 0:
            return "Infinite solutions"
        elif var_num == 0:  # 系数为 0 则无解
            return "No solution"
        else:
            return f"x={const_num // var_num}"
# @lc code=end
