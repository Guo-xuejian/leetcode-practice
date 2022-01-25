#
# @lc app=leetcode.cn id=1252 lang=python3
#
# [1252] 奇数值单元格的数目
#

# @lc code=start
class Solution:
    def oddCells(self, m: int, n: int, indices: List[List[int]]) -> int:
        matrix = [[0] * n for _ in range(m)]
        for x, y in indices:
            for i in range(m):
                matrix[i][y] += 1
            for i in range(n):
                matrix[x][i] += 1
        return sum(elem % 2 == 1 for line in matrix for elem in line)
# @lc code=end

