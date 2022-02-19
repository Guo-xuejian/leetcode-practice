#
# @lc app=leetcode.cn id=1926 lang=python3
#
# [1926] 迷宫中离入口最近的出口
#

# @lc code=start
class Solution:
    def nearestExit(self, maze: List[List[str]], entrance: List[int]) -> int:
        m, n = len(maze), len(maze[0])
        q = [(entrance[0], entrance[1])]
        directs = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def in_area(x, y):
            if 0 <= x < m and 0 <= y < n:
                return True
            return False

        maze[entrance[0]][entrance[1]] = "+"    # 起点已访问
        res = 1
        while q:
            for _ in range(len(q)):
                point = q[0]
                q = q[1:]
                for direct in directs:
                    x, y = point[0] + direct[0], point[1] + direct[1]
                    if in_area(x, y) and maze[x][y] == '.':
                        if x == 0 or y == 0 or x == m - 1 or y == n - 1:
                            return res
                        maze[x][y] = '+'    # 已访问
                        q.append((x, y))
            res += 1
        return -1
# @lc code=end
