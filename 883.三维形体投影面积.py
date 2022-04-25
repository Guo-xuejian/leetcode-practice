#
# @lc app=leetcode.cn id=883 lang=python3
#
# [883] 三维形体投影面积
#

# @lc code=start
class Solution:
    def projectionArea(self, grid: List[List[int]]) -> int:
        columnMax = []
        rowMax = [max(i) for i in grid]
        counts = 0
        for k in range(len(grid[0])):
            a = []
            for j in range(len(grid)):
                a.append(grid[j][k])
                if grid[j][k] == 0:
                    counts += 1
            columnMax.append(max(a))
        return sum(columnMax) + sum(rowMax) + len(grid)**2 - counts
# @lc code=end

