#
# @lc app=leetcode.cn id=6 lang=python3
#
# [6] Z 字形变换
#

# @lc code=start
class Solution:
    def convert(self, s: str, numRows: int) -> str:
        # 一行直接就可以返回
        if len(s) < 2:
            return s
        res = ["" for _ in range(numRows)]
        # row 对应行号,flag 记录当前是处于数组 s 的上升状态还是下降状态
        row = 0
        flag = True
        for i in range(len(s)):
            # 下降状态
            if flag:
                res[row] += s[i]
                # 下降至边界需要判定是否越界
                if row + 1 > numRows - 1:
                    # 越界则进入 上升状态,同时行号上升 -1
                    flag = False
                    row -= 1
                else:
                    # 不越界继续下降
                    row += 1
            # 上升状态
            else:
                res[row] += s[i]
                # 上升至边界需要判定是否越界
                if row - 1 < 0:
                    # 越界则进入 下降状态,同时行号下降 +1
                    flag = True
                    row += 1
                else:
                    # 不越界继续上升
                    row -= 1
        return "".join(res)
# @lc code=end

