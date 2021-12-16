#
# @lc app=leetcode.cn id=213 lang=python3
#
# [213] 打家劫舍 II
#

# @lc code=start
class Solution:
    def rob(self, nums: List[int]) -> int:
        length = len(nums)
        if length == 1:
            return nums[0]
        elif length == 2:
            return max(nums[0], nums[1])
        # 实际上和 打家劫舍 区别不大，只需要限制一下首尾相连
        # 最简单的方式就是两个取最大，一个不给尾，一个不给头
        return max(self.robRange(nums[:length-1]), self.robRange(nums[1:]))

    def robRange(self, numsNow):    # 打家劫舍的解法
        least, best = numsNow[0], max(numsNow[0], numsNow[1])
        for i in range(2, len(numsNow)):
            least, best = best, max(least+numsNow[i], best)
        return best
# @lc code=end
