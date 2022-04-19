#
# @lc app=leetcode.cn id=75 lang=python3
#
# [75] 颜色分类
#

# @lc code=start
class Solution:
    def sortColors(self, nums: List[int]) -> None:
        """
        Do not return anything, modify nums in-place instead.
        """
        # nums.sort(key = lambda e: e)
        def swap2Position(nums, target, index):
            for i in range(index, len(nums)):
                if nums[i] == target:
                    nums[index], nums[i] = nums[i], nums[index]
                    index += 1
            return index
        index = swap2Position(nums, 0, 0)  # 先移动 0
        swap2Position(nums, 1, index)  # 再移动 1， 2 自然就在最后连
        # 不能swap2Position(nums[index:], 1) 因为传递过去的会是一个拷贝
# @lc code=end

