#
# @lc app=leetcode.cn id=852 lang=python3
#
# [852] 山脉数组的峰顶索引
#

# @lc code=start
class Solution:
    def peakIndexInMountainArray(self, arr: List[int]) -> int:
        m = len(arr)
        left, right = 0, m - 1
        while left <= right:
            mid = (right - left) // 2 + left
            if arr[mid] > arr[mid - 1] and arr[mid] > arr[mid + 1]:
                return mid
            elif arr[mid] > arr[mid + 1]:
                right = mid
            else:
                left = mid
        return -1
# @lc code=end

