#
# @lc app=leetcode.cn id=1380 lang=python3
#
# [1380] 矩阵中的幸运数
#

# @lc code=start
class Solution:
    def luckyNumbers (self, matrix: List[List[int]]) -> List[int]:
        res = []
        m, n = len(matrix), len(matrix[0])
        row_min = [0 for _ in range(m)]
        col_max = [0 for _ in range(n)]

        for i, row in enumerate(matrix):
            row_min[i] = row[0]
            for j, x in enumerate(row):
                row_min[i] = min(row_min[i], x)
                col_max[j] = max(col_max[j], x)
        
        for i, row in enumerate(matrix):
            for j, x in enumerate(row):
                if x == row_min[i] and x == col_max[j]:
                    res.append(x)
        return res
# @lc code=end

