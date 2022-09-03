#
# @lc app=leetcode.cn id=646 lang=python3
#
# [646] 最长数对链
#

# @lc code=start
class Solution:
    def findLongestChain(self, pairs: List[List[int]]) -> int:
        pairs.sort()
        dp = [1] * len(pairs)
        for i in range(len(pairs)):
            for j in range(i):
                if pairs[i][0] > pairs[j][1]:
                    dp[i] = max(dp[i], dp[j] + 1)
        return dp[-1]
# @lc code=end

