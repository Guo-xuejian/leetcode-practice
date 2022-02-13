#
# @lc app=leetcode.cn id=200 lang=python3
#
# [200] 岛屿数量
#

# @lc code=start
class Solution:
    def numIslands(self, grid: List[List[str]]) -> int:
        m, n = len(grid), len(grid[0])
        # 上下左右
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def dfs(x, y):
            # 先判定越界，后判定是否是个还没有访问的陆地格子
            if x < 0 or y < 0 or x >= m or y >= n or grid[x][y] != "1":
                return
            grid[x][y] = "2"    # 剪枝，后续外层遍历不再访问
            for direction in directions:
                dfs(x + direction[0], y + direction[1])

        res = 0
        for i in range(m):
            for j in range(n):
                if grid[i][j] == "1":   # 只走还没有被访问的陆地格子
                    dfs(i, j)
                    res += 1            # dfs()内部只做访问和剪枝，外部才做结果记录
        return res
# @lc code=end
