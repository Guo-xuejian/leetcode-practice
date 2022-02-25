#
# @lc app=leetcode.cn id=283 lang=python3
#
# [283] 移动零
#

# @lc code=start
class Solution:
    def moveZeroes(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # def bubbleSortForTarget(nums, target):
        #     for i in range(len(nums)):
        #         for j in range(len(nums) - i - 1):
        #             if nums[j] == target:
        #                 nums[j], nums[j + 1] = nums[j + 1], nums[j]
        
        # bubbleSortForTarget(nums, 0)
        n = len(nums)
        left = right = 0
        # 将所有的 0 聚集在一起，同步向后移，不是每个 0 都移动到最后，所有的 0 聚在一起在末尾即可
        while right < n:
            if nums[right] != 0:
                nums[left], nums[right] = nums[right], nums[left]
                left += 1
            right += 1
# @lc code=end

