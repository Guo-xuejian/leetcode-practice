class Solution:
    def maxValue(self, grid: List[List[int]]) -> int:
        for i in range(len(grid)):
            for j in range(len(grid[0])):
                # 如果是初始位置，也就是没有左边和上边，那么跳过，因为无法给出运算————为了保证转移方程的完整性
                if i == 0 and j == 0:
                    continue
                # 第一行，没有上边，只能往左边加
                if i == 0:
                    grid[i][j] += grid[i][j - 1]
                # 第一列，没有左边，只能往上边加
                elif j == 0:
                    grid[i][j] += grid[i - 1][j]
                # 不是墙壁（特指左边和上边）
                else:
                    grid[i][j] += max(grid[i - 1][j], grid[i][j - 1])
        return grid[-1][-1]