#
# @lc app=leetcode.cn id=566 lang=python3
#
# [566] 重塑矩阵
#

# @lc code=start
class Solution:
    def matrixReshape(self, mat: List[List[int]], r: int, c: int) -> List[List[int]]:
        m, n = len(mat), len(mat[0])
        if m * n != r * c:
            return mat
        
        res = [[0] * c for _ in range(r)]
        for x in range(m * n):
            res[x // c][x % c] = mat[x // n][x % n]
        
        return res
# @lc code=end

