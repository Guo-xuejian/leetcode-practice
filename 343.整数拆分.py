#
# @lc app=leetcode.cn id=343 lang=python3
#
# [343] 整数拆分
#

# @lc code=start
class Solution:
    def integerBreak(self, n: int) -> int:
        # You're making it pretty complicated.

        # If an optimal product contains a factor f >= 4, then you can replace it with factors 2 and f-2 without losing optimality, as 2*(f-2) = 2f-4 >= f. So you never need a factor greater than or equal to 4, meaning you only need factors 1, 2 and 3 (and 1 is of course wasteful and you'd only use it for n=2 and n=3, where it's needed).

        # For the rest I agree, 3*3 is simply better than 2*2*2, so you'd never use 2 more than twice.
        if n <= 3:
            return n-1
        ans = 1
        while n > 4:    # 无限拆出 3 相乘，完美符合总和要求
            ans *= 3
            n -= 3
        return ans * n  # 最后乘上剩下的数字
# @lc code=end

