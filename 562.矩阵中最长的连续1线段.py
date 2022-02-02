# 562. 矩阵中最长的连续1线段
# 给定一个 m x n 的二进制矩阵 mat ，返回矩阵中最长的连续1线段。

# 这条线段可以是水平的、垂直的、对角线的或者反对角线的。

 

# 示例 1:



# 输入: mat = [[0,1,1,0],[0,1,1,0],[0,0,0,1]]
# 输出: 3
# 示例 2:



# 输入: mat = [[1,1,1,1],[0,1,1,0],[0,0,0,1]]
# 输出: 4
 

# 提示:

# m == mat.length
# n == mat[i].length
# 1 <= m, n <= 104
# 1 <= m * n <= 104
# mat[i][j] 不是 0 就是 1.


class Solution:
    def longestLine(self, mat: List[List[int]]) -> int:
        m, n = len(mat), len(mat[0])

        # n + 2, m + 2 多出来的 2 是首尾边界的额外值，这样可以省去很多边界判定的事情
        dp = [[[0] * 4 for _ in range(n+2)] for _ in range(m+2)]

        res = 0
        # 由于多申请空间，于是从 1 开始，同时 mat 访问需要减一
        for i in range(1, m + 1):
            for j in range(1, n + 1):
                if mat[i-1][j-1] == 1:
                    dp[i][j][0] = dp[i - 1][j][0] + 1       # 水平
                    dp[i][j][1] = dp[i][j - 1][1] + 1       # 垂直
                    dp[i][j][2] = dp[i - 1][j - 1][2] + 1   # 对角
                    dp[i][j][3] = dp[i - 1][j + 1][3] + 1   # 反对角
                    res = max(res, max(dp[i][j]))
        return res