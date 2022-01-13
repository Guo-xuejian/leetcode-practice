#
# @lc app=leetcode.cn id=910 lang=python3
#
# [910] 最小差值 II
#

# @lc code=start
class Solution:
    def smallestRangeII(self, nums: List[int], k: int) -> int:
        nums.sort()
        length = len(nums)
        res = nums[length - 1] - nums[0]
        # 最差的情况==最小的加上k和最大的减去k
        p1 = nums[0] + k
        p4 = nums[length-1] - k
        # 有序，所以较小的加k 较大的-k 比较
        for i in range(length-1):
            p2 = nums[i] + k
            p3 = nums[i+1] - k
            h = max(p2, p4)
            l = min(p1, p3)
            res = min(res, h-l)
        return res
# @lc code=end
