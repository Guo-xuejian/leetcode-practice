#
# @lc app=leetcode.cn id=67 lang=python3
#
# [67] 二进制求和
#

# @lc code=start
class Solution:
    def addBinary(self, a: str, b: str) -> str:
        add_status = False
        m, n = len(a), len(b)
        res = []
        index1, index2 = m - 1, n - 1
        while index1 >= 0 or index2 >= 0:
            num1, num2 = 0, 0
            if index1 > -1:
                num1 = int(a[index1])
                index1 -= 1
            if index2 > -1:
                num2 = int(b[index2])
                index2 -= 1

            sum_now = num1 + num2 + int(add_status)
            if sum_now >= 2:
                add_status = True
                res.append(str(sum_now % 2))
            else:
                add_status = False
                res.append(str(sum_now))
        if add_status:
            res.append("1")
        return "".join(res[::-1])
# @lc code=end
