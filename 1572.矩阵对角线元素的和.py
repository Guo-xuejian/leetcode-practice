#
# @lc app=leetcode.cn id=1572 lang=python3
#
# [1572] 矩阵对角线元素的和
#

# @lc code=start
class Solution:
    def diagonalSum(self, mat: List[List[int]]) -> int:
        m = len(mat)
        index = 0
        res = 0
        while index < m:
            res += mat[index][index] + mat[index][m - 1 - index] + mat[m - 1 - index][index] + mat[m - 1 - index][m - 1 - index]
            index += 1
        res //= 2
        if m % 2 != 0:
            res -= mat[m//2][m//2]
        return res
# @lc code=end

