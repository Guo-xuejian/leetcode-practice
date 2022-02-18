#
# @lc app=leetcode.cn id=417 lang=python3
#
# [417] 太平洋大西洋水流问题
#

# @lc code=start
class Solution:
    def pacificAtlantic(self, heights: List[List[int]]) -> List[List[int]]:
        res = []
        m, n = len(heights), len(heights[0])
        visited = [[0 for _ in range(n)] for _ in range(m)]
        pacific = [[0 for _ in range(n)] for _ in range(m)]
        atlantic = [[0 for _ in range(n)] for _ in range(m)]
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def in_area(x, y):
            if 0 <= x < m and 0 <= y < n:
                return True
            return False

        def dfs(x, y, flag):
            if visited[x][y] == 1:
                return
            visited[x][y] = 1
            if flag:
                pacific[x][y] = 1
            else:
                atlantic[x][y] = 1
            for i in range(4):
                curr_x, curr_y = x + directions[i][0], y + directions[i][1]
                if not in_area(curr_x, curr_y) or heights[x][y] > heights[curr_x][curr_y]:
                    continue
                dfs(curr_x, curr_y, flag)
            return

        for i in range(n):
            dfs(0, i, True)
        visited = [[0 for _ in range(n)] for _ in range(m)]
        for i in range(m):
            dfs(i, 0, True)
        visited = [[0 for _ in range(n)] for _ in range(m)]
        for i in range(n):
            dfs(m - 1, i, False)
        visited = [[0 for _ in range(n)] for _ in range(m)]
        for i in range(m):
            dfs(i, n - 1, False)

        for i in range(m):
            for j in range(n):
                if pacific[i][j] == 1 and atlantic[i][j] == 1:
                    res.append([i, j])
        return res
# @lc code=end
