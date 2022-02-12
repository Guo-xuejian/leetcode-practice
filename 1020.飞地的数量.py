#
# @lc app=leetcode.cn id=1020 lang=python3
#
# [1020] 飞地的数量
#

# @lc code=start
class Solution:
    def numEnclaves(self, grid: List[List[int]]) -> int:
        # 不是飞地有两种情况，一种常规在内部可以连接至边界，一种是就在边界
        # 那么就从边界开始遍历，深度优先搜索，遍历出所有满足条件的陆地格子，最后内部剩下的没有被访问的陆地格子就是飞地
        m, n = len(grid), len(grid[0])
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def dfs(i, j):
            # 索引越界，海洋格子，或者已经被访问过的，都不再继续
            if i < 0 or i >= m or j < 0 or j >= n or grid[i][j] == 0 or grid[i][j] == 3:
                return
            grid[i][j] = 3
            for x, y in directions:
                dfs(i + x, j + y)

        for i in range(m):
            dfs(i, 0)       # 每一层的左边界
            dfs(i, n - 1)   # 右边界
        for j in range(1, n - 1):
            dfs(0, j)       # 每一列的上边界
            dfs(m - 1, j)   # 下边界
        # 由于处于边界的陆地格子不需要考虑，所以只判定内部格子
        res = 0
        for i in range(1, m - 1):
            for j in range(1, n - 1):
                if grid[i][j] == 1:
                    res += 1
        return res
# @lc code=end
