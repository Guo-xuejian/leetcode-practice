#
# @lc app=leetcode.cn id=2028 lang=python3
#
# [2028] 找出缺失的观测数据
#

# @lc code=start
class Solution:
    def missingRolls(self, rolls: List[int], mean: int, n: int) -> List[int]:
        missingSum = mean * (n + len(rolls)) - sum(rolls)
        if not n <= missingSum <= n * 6:
            return []
        quotient, remainder = divmod(missingSum, n)
        return [quotient + 1] * remainder + [quotient] * (n - remainder)

# @lc code=end

