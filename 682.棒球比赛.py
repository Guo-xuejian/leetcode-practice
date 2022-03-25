#
# @lc app=leetcode.cn id=682 lang=python3
#
# [682] 棒球比赛
#

# @lc code=start
class Solution:
    def calPoints(self, ops: List[str]) -> int:
        arr = []
        res = 0
        for one in ops:
            if one == "+":
                arr.append(arr[-1] + arr[-2])
            elif one == "D":
                arr.append(arr[-1] * 2)
            elif one == "C":
                res -= arr.pop()
                continue
            else:
                arr.append(int(one))
            res += arr[-1]
        return res
# @lc code=end

