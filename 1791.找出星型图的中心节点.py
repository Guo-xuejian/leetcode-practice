#
# @lc app=leetcode.cn id=1791 lang=python3
#
# [1791] 找出星型图的中心节点
#

# @lc code=start
class Solution:
    def findCenter(self, edges: List[List[int]]) -> int:
        m = len(edges) + 1
        temp = [0 for _ in range(m)]
        for edge in edges:
            temp[edge[0] - 1] += 1
            temp[edge[1] - 1] += 1
        
        res = 0
        for i in range(m):
            if temp[i] == m - 1:
                res = i + 1
                break
        return res
# @lc code=end

