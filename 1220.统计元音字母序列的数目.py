#
# @lc app=leetcode.cn id=1220 lang=python3
#
# [1220] 统计元音字母序列的数目
#

# @lc code=start
class Solution:
    def countVowelPermutation(self, n: int) -> int:
        mod_int = 1e9 + 7
        dp = [1, 1, 1, 1, 1]
        for i in range(1, n):
            a_end, e_end, i_end, o_end, u_end = dp[0], dp[1], dp[2], dp[3], dp[4]
            dp[0] = (e_end + i_end + u_end) % mod_int
            dp[1] = (a_end + i_end) % mod_int
            dp[2] = (e_end + o_end) % mod_int
            dp[3] = i_end % mod_int
            dp[4] = (i_end + o_end) % mod_int
        return int(sum(dp) % mod_int)
# @lc code=end
