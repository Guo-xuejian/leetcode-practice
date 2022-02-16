#
# @lc app=leetcode.cn id=1162 lang=python3
#
# [1162] 地图分析
#

# @lc code=start
class Solution:
    def maxDistance(self, grid: List[List[int]]) -> int:
        m, n = len(grid), len(grid[0])
        queue = []
        for i in range(m):
            for j in range(n):
                if grid[i][j] == 1:
                    queue.append((i, j))
        
        if len(queue) == 0 or len(queue) == m * n:
            return -1
        
        res = 0
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        while queue:
            res += 1
            temp_queue = queue
            queue = []
            while temp_queue:
                curr = temp_queue[0]
                temp_queue = temp_queue[1:]
                for i in range(4):
                    x = curr[0] + directions[i][0]
                    y = curr[1] + directions[i][1]
                    if x < 0 or x >= m or y < 0 or y >= n or grid[x][y] != 0:
                        continue
                    queue.append((x, y))
                    grid[x][y] = 2
        return res - 1
# @lc code=end

