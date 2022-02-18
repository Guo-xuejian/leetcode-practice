#
# @lc app=leetcode.cn id=542 lang=python3
#
# [542] 01 矩阵
#

# @lc code=start
class Solution:
    def updateMatrix(self, mat: List[List[int]]) -> List[List[int]]:
        m, n = len(mat), len(mat[0])
        q = []
        for i in range(m):
            for j in range(n):
                if mat[i][j] == 0:
                    q.append((i, j))
                else:
                    mat[i][j] = -1
        
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        while len(q) > 0:
            point = q[0]
            q = q[1:]
            for direction in directions:
                x = point[0] + direction[0]
                y = point[1] + direction[1]

                if 0 <= x < m and 0 <= y < n and mat[x][y] == -1:
                    mat[x][y] = mat[point[0]][point[1]] + 1
                    q.append((x, y))
        return mat
# @lc code=end

