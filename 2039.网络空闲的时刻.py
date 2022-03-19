#
# @lc app=leetcode.cn id=2039 lang=python3
#
# [2039] 网络空闲的时刻
#

# @lc code=start
class Solution:
    def networkBecomesIdle(self, edges: List[List[int]], patience: List[int]) -> int:
        n = len(patience)
        g = [[] for _ in range(n)]
        for u, v in edges:
            g[u].append(v)
            g[v].append(u)

        vis = [False] * n
        vis[0] = True
        q = deque([0])
        ans, dist = 0, 1
        while q:
            for _ in range(len(q)):
                u = q.popleft()
                for v in g[u]:
                    if vis[v]:
                        continue
                    vis[v] = True
                    q.append(v)
                    ans = max(ans, (dist * 2 - 1) // patience[v] * patience[v] + dist * 2 + 1)
            dist += 1
        return ans

# @lc code=end

