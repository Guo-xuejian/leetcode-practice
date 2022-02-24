#
# @lc app=leetcode.cn id=374 lang=python3
#
# [374] 猜数字大小
#

# @lc code=start
# The guess API is already defined for you.
# @param num, your guess
# @return -1 if my number is lower, 1 if my number is higher, otherwise return 0
# def guess(num: int) -> int:

class Solution:
    def guessNumber(self, n: int) -> int:
        left, right = 0, n
        while left <= right:
            mid = (left + right) // 2
            guess_result = guess(mid)
            if guess_result == 0:
                return mid
            elif guess_result == 1:
                left = mid + 1
            else:
                right = mid - 1
        return -1
        
# @lc code=end

