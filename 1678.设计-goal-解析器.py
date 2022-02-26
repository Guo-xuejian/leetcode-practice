#
# @lc app=leetcode.cn id=1678 lang=python3
#
# [1678] 设计 Goal 解析器
#

# @lc code=start
class Solution:
    def interpret(self, command: str) -> str:
        # return command.replace("()", "o").replace("(al)", "al")
        pre_letter = command[0]
        res = []
        index = 1
        if pre_letter == "G":
            res.append("G")
        while index < len(command):
            if command[index] == ")" and pre_letter == "(":
                res.append("o")
            elif command[index] == "a":
                pre_letter = command[index]
                index += 3
                res.append("al")
                continue
            elif command[index] == "G":
                res.append("G")
            pre_letter = command[index]
            index += 1
        return "".join(res)
# @lc code=end
