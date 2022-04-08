#
# @lc app=leetcode.cn id=1855 lang=python3
#
# [1855] 下标对中的最大距离
#

# @lc code=start
class Solution:
    def maxDistance(self, nums1: List[int], nums2: List[int]) -> int:
        n1, n2 = len(nums1), len(nums2)
        i = 0
        res = 0
        for j in range(n2):
            while i < n1 and nums1[i] > nums2[j]:
                i += 1
            if i < n1:
                res = max(res, j - i)
        return res
# @lc code=end

