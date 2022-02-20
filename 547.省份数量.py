#
# @lc app=leetcode.cn id=547 lang=python3
#
# [547] 省份数量
#

# @lc code=start
class Solution:
    def findCircleNum(self, isConnected: List[List[int]]) -> int:
        city = len(isConnected)
        visited = {}
        circles = 0
        
        def dfs(i):
            for j in range(city):
                if isConnected[i][j] == 1 and j not in visited:
                    visited[j] = True
                    dfs(j)
        
        for i in range(city):
            if i not in visited:
                dfs(i)
                circles += 1
        return circles
# @lc code=end

