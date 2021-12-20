#
# @lc app=leetcode.cn id=1887 lang=python3
#
# [1887] 使数组元素相等的减少操作次数
#

# @lc code=start
class Solution:
    def reductionOperations(self, nums: List[int]) -> int:
        nums.sort()
        length = len(nums)
        res = 0     # 总操作次数
        curr = 0    # 每个元素操作次数
        for i in range(1, length):
            if nums[i] != nums[i - 1]:  # 不重复则需要进行数字变换
                curr += 1
            res += curr     # 由于是渐进式修改，所有每次都需要加上新修改的次数
        return res
# @lc code=end

