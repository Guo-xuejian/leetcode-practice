#
# @lc app=leetcode.cn id=997 lang=python3
#
# [997] 找到小镇的法官
#

# @lc code=start
class Solution:
    def findJudge(self, n: int, trust: List[List[int]]) -> int:
        res = [0 for _ in range(n)]
        for one in trust:
            res[one[1]-1] += 1      # 被相信者加一
            res[one[0]-1] -= 1      # 相信者减一
        for i in range(n):
            if res[i] == n - 1:     # 满足所有热都相信他
                return i + 1
        return -1
# @lc code=end

