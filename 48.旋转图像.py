#
# @lc app=leetcode.cn id=48 lang=python3
#
# [48] 旋转图像
#

# @lc code=start
class Solution:
    def rotate(self, matrix: List[List[int]]) -> None:
        """
        Do not return anything, modify matrix in-place instead.
        """
        n = len(matrix)
        temp = [[0 for _ in range(n)] for _ in range(n)]
        for i in range(n):
            for j in range(n):
                temp[j][n - i - 1] = matrix[i][j]
        matrix[:] = temp
# @lc code=end

