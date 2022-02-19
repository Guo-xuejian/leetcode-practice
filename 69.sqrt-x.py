#
# @lc app=leetcode.cn id=69 lang=python3
#
# [69] Sqrt(x)
#

# @lc code=start
class Solution:
    def mySqrt(self, x: int) -> int:
        res, left, right = 0, 0, x
        while left <= right:
            mid = (left + right) // 2
            if mid * mid <= x:
                res = mid
                left = mid + 1
            else:
                right = mid - 1
        return res
# @lc code=end

