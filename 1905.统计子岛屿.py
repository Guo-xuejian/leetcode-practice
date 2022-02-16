#
# @lc app=leetcode.cn id=1905 lang=python3
#
# [1905] 统计子岛屿
#

# @lc code=start
class Solution:
    def countSubIslands(self, grid1: List[List[int]], grid2: List[List[int]]) -> int:
        # 得到矩阵的行和列
        m, n = len(grid1), len(grid1[0])
        directions = [(-1, 0), (1, 0), (0, -1), (0, 1)]

        # dfs，搜索一片岛屿
        def dfs(x, y):
            if not (0 <= x < m and 0 <= y < n and grid2[x][y] == 1):
                return
            # 将搜索过的岛屿置0，防止重复搜索
            grid2[x][y] = 0

            for direction in directions:
                dfs(x + direction[0], y + direction[1])

        # 首先将 2 中岛屿不被 1 中包含的剔除掉
        for i in range(m):
            for j in range(n):
                if grid2[i][j] == 1 and grid1[i][j] == 0:
                    dfs(i, j)

        # 计算子岛屿数量
        res = 0
        for i in range(m):
            for j in range(n):
                if grid2[i][j] == 1:
                    res += 1
                    dfs(i, j)
        return res
# @lc code=end
