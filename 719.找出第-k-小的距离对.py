#
# @lc app=leetcode.cn id=719 lang=python3
#
# [719] 找出第 k 小的距离对
#

# @lc code=start
class Solution:
    def smallestDistancePair(self, nums: List[int], k: int) -> int:
        def count(mid: int) -> int:
            cnt = 0
            for j, num in enumerate(nums):
                i = bisect_left(nums, num - mid, 0, j)
                cnt += j - i
            return cnt

        nums.sort()
        return bisect_left(range(nums[-1] - nums[0]), k, key=count)
# @lc code=end

