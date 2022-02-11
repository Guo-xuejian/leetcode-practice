#
# @lc app=leetcode.cn id=640 lang=python3
#
# [640] 求解方程
#

# @lc code=start
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
