#
# @lc app=leetcode.cn id=198 lang=python3
#
# [198] 打家劫舍
#

# @lc code=start
class Solution:
    def rob(self, nums: List[int]) -> int:
        length = len(nums)
        if length == 1:
            return nums[0]
        elif length == 2:
            return max(nums[0], nums[1])
        # least 和 best 相等的情况下是因为 best 两边加起来没有 best 大
        # 所以在判时下一个索引对应最大值虽然和 best 左边元素组合，但是计算后最大值仍然是 best
        least, best = nums[0], max(nums[0], nums[1])
        for i in range(2, length):
            least, best = best, max(least + nums[i], best)
        return best
# @lc code=end
