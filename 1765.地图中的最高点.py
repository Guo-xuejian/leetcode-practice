#
# @lc app=leetcode.cn id=1765 lang=python3
#
# [1765] 地图中的最高点
#

# @lc code=start
from collections import deque


class Solution:
    def highestPeak(self, isWater: List[List[int]]) -> List[List[int]]:
        m, n  = len(isWater), len(isWater[0])
        res = [[0] * n for _ in range(m)]
        d = deque()
        for i in range(m):
            for j in range(n):
                if isWater[i][j]:
                    d.append((i,j))
                res[i][j] = 0 if isWater[i][j] else - 1
        
        dirs = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        h = 1
        while d:
            size = len(d)
            for _ in range(size):
                x, y = d.popleft()
                for di in dirs:
                    nx, ny = x + di[0], y + di[1]
                    if 0 <= nx < m and 0 <= ny < n and res[nx][ny] == -1:
                        res[nx][ny] = h
                        d.append((nx, ny))
            h += 1
        return res
# @lc code=end

