#
# @lc app=leetcode.cn id=1550 lang=python3
#
# [1550] 存在连续三个奇数的数组
#

# @lc code=start
class Solution:
    def threeConsecutiveOdds(self, arr: List[int]) -> bool:
        def is_odd(num):
            return num % 2 == 1
        three_limited = 0
        for num in arr:
            if is_odd(num):
                three_limited += 1
            else:
                three_limited = 0
            if three_limited == 3:
                return True
        return False
# @lc code=end

