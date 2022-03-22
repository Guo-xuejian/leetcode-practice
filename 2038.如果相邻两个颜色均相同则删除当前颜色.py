#
# @lc app=leetcode.cn id=2038 lang=python3
#
# [2038] 如果相邻两个颜色均相同则删除当前颜色
#

# @lc code=start
class Solution:
    def winnerOfGame(self, colors: str) -> bool:
        res = 0
        for i in range(1, len(colors)-1):
            if colors[i] == colors[i - 1] and colors[i] == colors[i + 1]:
               res += 1 if colors[i] == "A" else -1
        return res > 0
# @lc code=end

