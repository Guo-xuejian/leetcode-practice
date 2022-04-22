#
# @lc app=leetcode.cn id=587 lang=python3
#
# [587] 安装栅栏
#

# @lc code=start
class Solution:
    def outerTrees(self, trees: List[List[int]]) -> List[List[int]]:
        n = len(trees)
        if n < 3: return trees
        trees.sort(key = lambda x:(x[0],x[1]))
        cross = lambda a,b,c:(b[0]-a[0])*(c[1]-b[1])-(b[1]-a[1])*(c[0]-b[0])        
        lower = []
        for p in trees:
            while len(lower) > 1 and cross(lower[-2],lower[-1],p) < 0: lower.pop()
            lower.append((p[0],p[1]))        
        upper = []
        for p in reversed(trees):
            while len(upper) > 1 and cross(upper[-2],upper[-1],p) < 0: upper.pop()
            upper.append((p[0],p[1]))            
        return list(set(lower[:-1] + upper[:-1]))
# @lc code=end

