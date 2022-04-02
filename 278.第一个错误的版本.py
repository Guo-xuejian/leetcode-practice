#
# @lc app=leetcode.cn id=278 lang=python3
#
# [278] 第一个错误的版本
#

# @lc code=start
# The isBadVersion API is already defined for you.
# @param version, an integer
# @return an integer
# def isBadVersion(version):

class Solution:
    def firstBadVersion(self, n):
        """
        :type n: int
        :rtype: int
        """
        res = n
        left, right = 0, n
        while left <= right:
            mid = left + (right - left) // 2
            if isBadVersion(mid):
                res = mid
                right = mid - 1
            else:
                left = mid + 1
        return res
        
# @lc code=end

