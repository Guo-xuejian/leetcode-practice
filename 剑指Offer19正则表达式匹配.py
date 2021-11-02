class Solution:
    def isMatch(self, s: str, p: str) -> bool:
        m, n = len(s) + 1, len(p) + 1
        # 创建动态规划矩阵
        dp = [[False] * n for _ in range(m)]
        dp[0][0] = True

        for j in range(2, n, 2):
            dp[0][j] = dp[0][j - 2] and p[j - 1] == '*'
        
        for i in range(1, m):
            for j in range(1, n):
                if p[j - 1] == '*':
                    if dp[i][j - 2]:
                        dp[i][j] = True
                    elif dp[i - 1][j] and s[i - 1] == p[j - 2]:
                        dp[i][j] = True
                    elif dp[i - 1][j] and p[j - 2] == '.':
                        dp[i][j] = True
                else:
                    if dp[i - 1][j - 1] and s[i - 1] == p[j - 1]:
                        dp[i][j] = True
                    elif dp[i - 1][j - 1] and p[j - 1] == '.':
                        dp[i][j] = True
        return dp[-1][-1]