#
# @lc app=leetcode.cn id=934 lang=python3
#
# [934] 最短的桥
#

# @lc code=start
class Solution:
    def shortestBridge(self, grid: List[List[int]]) -> int:
        q = []
        m, n  = len(grid), len(grid[0])
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        
        def dfs(x, y):
            if not (0 <= x < m and 0 <= y < n and grid[x][y] == 1):
                return
            grid[x][y] = 0
            q.append((x, y))
            for direction in directions:
                dfs(x + direction[0], y + direction[1])
        
        visited = set()
        res = 0
        find = False
        for i in range(m):
            for j in range(n):
                if grid[i][j] and not find:
                    find = True
                    dfs(i, j)
        while q:
            for _ in range(len(q)):
                x, y = q[0]
                q = q[1:]
                for direction in directions:
                    curr_x, curr_y = x +direction[0], y + direction[1]

                    if not (0 <= curr_x < m and 0 <= curr_y < n) or (curr_x, curr_y) in visited:
                        continue
                    if grid[curr_x][curr_y]:
                        return res
                    else:
                        q.append((curr_x, curr_y))
                        visited.add((curr_x, curr_y))
            res += 1
        return res
# @lc code=end

