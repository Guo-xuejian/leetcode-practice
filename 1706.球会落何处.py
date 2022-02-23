#
# @lc app=leetcode.cn id=1706 lang=python3
#
# [1706] 球会落何处
#

# @lc code=start
class Solution:
    def findBall(self, grid: List[List[int]]) -> List[int]:
        m, n = len(grid), len(grid[0])
        def dfs(i, j):
            if i == m:
                return j
            if grid[i][j] == 1:
                if j == n - 1 or grid[i][j + 1] == -1:
                    return -1
                return dfs(i + 1, j + 1)
            else:
                if j == 0 or grid[i][j - 1] == 1:
                    return -1
                return dfs(i + 1, j - 1)
        return [dfs(0, j) for j in range(n)]
# @lc code=end

