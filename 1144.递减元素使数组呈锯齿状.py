#
# @lc app=leetcode.cn id=1144 lang=python3
#
# [1144] 递减元素使数组呈锯齿状
#

# @lc code=start
class Solution:
    def movesToMakeZigzag(self, nums: List[int]) -> int:
        res_odd, res_even = 0, 0
        length = len(nums)
        for i in range(length):
            # 需要注意，数组索引奇偶和常规数字不一样，从零开始，索引 1 为实际偶数为
            if i % 2 == 0:
                # 奇数需要大于相邻数字
                pre = nums[i] - nums[i - 1] + 1 if i > 0 and nums[i] >= nums[i - 1] else 0
                after = nums[i] - nums[i + 1] + 1 if i < length - 1 and nums[i] >= nums[i + 1] else 0
                # 同理，较大值
                res_odd += max(pre, after)
            else:
                # 或者偶数需要大于相邻数字
                pre = nums[i] - nums[i - 1] + 1 if i > 0 and nums[i] >= nums[i - 1] else 0
                after = nums[i] - nums[i + 1] + 1 if i < length - 1 and nums[i] >= nums[i + 1] else 0
                # 为满足两边，取较大值
                res_even += max(pre, after)
        return min(res_odd, res_even)
# @lc code=end

