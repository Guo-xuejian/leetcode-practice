#
# @lc app=leetcode.cn id=202 lang=python3
#
# [202] 快乐数
#

# @lc code=start
class Solution:
    def isHappy(self, n: int) -> bool:
        def step(n):
            res = 0
            while n > 0:
                res += (n % 10) * (n % 10)
                n //= 10
            return res
        
        visited = {}
        while n != 1 and n not in visited:
            visited[n] = True
            n = step(n)
        return n == 1
# @lc code=end

