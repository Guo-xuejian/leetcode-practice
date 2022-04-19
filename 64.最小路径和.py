#
# @lc app=leetcode.cn id=64 lang=python3
#
# [64] 最小路径和
#

# @lc code=start
class Solution:
    def minPathSum(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        temp = [[0] * n for _ in range(m)]
        for i in range(m):
            if i == 0:
                temp[0][:] = grid[0][:]
                continue
            temp[i][0] = grid[i][0]
        for i in range(1, n):
            temp[0][i] = grid[0][i] + temp[0][i - 1]
        for j in range(1, m):
            temp[j][0] = grid[j][0] + temp[j - 1][0]
        for i in range(1, m):
            for j in range(1, n):
                temp[i][j] = grid[i][j] + min(temp[i - 1][j], temp[i][j - 1])
        return temp[m - 1][n - 1]
# @lc code=end

