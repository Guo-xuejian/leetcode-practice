# 剑指 Offer 47. 礼物的最大价值
# 在一个 m*n 的棋盘的每一格都放有一个礼物，每个礼物都有一定的价值（价值大于 0）。你可以从棋盘的左上角开始拿格子里的礼物，并每次向右或者向下移动一格、直到到达棋盘的右下角。给定一个棋盘及其上面的礼物的价值，请计算你最多能拿到多少价值的礼物？

# 示例 1:

# 输入:
# [
#   [1,3,1],
#   [1,5,1],
#   [4,2,1]
# ]
# 输出: 12
# 解释: 路径 1→3→5→2→1 可以拿到最多价值的礼物

# 提示：

# 0 < grid.length <= 200
# 0 < grid[0].length <= 200


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

# 2022-04-20
# class Solution:
#     def maxValue(self, grid: List[List[int]]) -> int:
#         m, n = len(grid), len(grid[0])
#         for i in range(1, n):
#             grid[0][i] += grid[0][i - 1]
#         for j in range(1, m):
#             grid[j][0] += grid[j - 1][0]
#         for i in range(1, m):
#             for j in range(1, n):
#                 grid[i][j] += max(grid[i - 1][j], grid[i][j - 1])
#         return grid[m - 1][n - 1]