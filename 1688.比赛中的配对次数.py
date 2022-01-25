#
# @lc app=leetcode.cn id=1688 lang=python3
#
# [1688] 比赛中的配对次数
#

# @lc code=start
class Solution:
    def numberOfMatches(self, n: int) -> int:
        res = 0
        while n > 1:
            res += int(n/2)     # 加总配对次数
            if n % 2 != 0:      # 奇数除二后加一
                n = int(n/2) + 1
            else:               # 偶数直接除二
                n = int(n/2)
        return res
# @lc code=end

