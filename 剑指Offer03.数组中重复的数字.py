# 剑指 Offer 03. 数组中重复的数字
# 找出数组中重复的数字。


# 在一个长度为 n 的数组 nums 里的所有数字都在 0～n-1 的范围内。数组中某些数字是重复的，但不知道有几个数字重复了，也不知道每个数字重复了几次。请找出数组中任意一个重复的数字。

# 示例 1：

# 输入：
# [2, 3, 1, 0, 2, 5, 3]
# 输出：2 或 3 
 

# 限制：

# 2 <= n <= 100000

class Solution:
    def findRepeatNumber(self, nums: List[int]) -> int:
        nums.sort()
        for i in range(1, len(nums)):
            if nums[i] == nums[i - 1]:
                return nums[i]

# 2022-04-15
# class Solution:
#     def findRepeatNumber(self, nums: List[int]) -> int:
#         num_dict = {}
#         for num in nums:
#             if num not in num_dict:
#                 num_dict[num] = 1
#             else:
#                 return num
#         return -1