#
# @lc app=leetcode.cn id=31 lang=python3
#
# [31] 下一个排列
#

# @lc code=start
class Solution:
    def nextPermutation(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        if len(nums) <= 1:
            return
        
        i, j, k = len(nums)-2, len(nums)-1, len(nums)-1
        # find: A[i]<A[j]
        # 这里要寻找的就是从后往前第一个下降的（注意方向，后--前），那么在后续的reverse 步骤中 j--end也一定是递减的（前---后），对应也就是最小排列
        while i >= 0 and nums[i] >= nums[j]:
            i -= 1
            j -= 1
        
        if i >= 0:  # 不是最后一个排列
            # find: A[i]<A[k]，从后往前找第一个比A[i]大的，有可能越过 A[i]
            while nums[i] >= nums[k]:
                k -= 1
            # swap A[i], A[k]，交换完成下一个排列
            nums[i], nums[k] = nums[k], nums[i]
        
        # reverse A[j:end]，将后续递增改为递减，符合最小要求
        i, j = j, len(nums)-1
        while i < j:
            nums[i], nums[j] = nums[j], nums[i]
            i += 1
            j -= 1
# @lc code=end

