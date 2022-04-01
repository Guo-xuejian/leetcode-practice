#
# @lc app=leetcode.cn id=367 lang=python3
#
# [367] 有效的完全平方数
#

# @lc code=start
class Solution:
    def isPerfectSquare(self, num: int) -> bool:
        left, right = 0, num
        while left <= right:
            mid = (right - left) // 2 + left
            curr = mid * mid
            if curr == num:
                return True
            elif curr > num:
                right = mid - 1
            else:
                left = mid + 1
        return False
# @lc code=end

