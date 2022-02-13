#
# @lc app=leetcode.cn id=1254 lang=python3
#
# [1254] 统计封闭岛屿的数目
#

# @lc code=start
class Solution:
    def closedIsland(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        directions = [[1, 0], [-1, 0], [0, 1], [0, -1]]

        def dfs(x, y):
            if x < 0 or y < 0 or x >= m or y >= n or grid[x][y] == 1:
                return
            grid[x][y] = 1      # 外界已经计数，清除所有相连的陆地
            for direction in directions:
                dfs(x + direction[0], y + direction[1])

        for i in range(m):
            dfs(i, 0)
            dfs(i, n - 1)

        for j in range(n):
            dfs(0, j)
            dfs(m - 1, j)

        res = 0
        for i in range(1, m - 1):
            for j in range(1, n - 1):
                if grid[i][j] == 0:
                    res += 1
                    dfs(i, j)
        return res
# @lc code=end
