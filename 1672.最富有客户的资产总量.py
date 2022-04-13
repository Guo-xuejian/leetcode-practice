#
# @lc app=leetcode.cn id=1672 lang=python3
#
# [1672] 最富有客户的资产总量
#

# @lc code=start
class Solution:
    def maximumWealth(self, accounts: List[List[int]]) -> int:
        res = 0
        for account in accounts:
            res = max(res, sum(account))
        return res
# @lc code=end

