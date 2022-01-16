#
# @lc app=leetcode.cn id=1985 lang=python3
#
# [1985] 找出数组中的第 K 大整数
#

# @lc code=start
class Solution:
    def kthLargestNumber(self, nums: List[str], k: int) -> str:
        # 无赖做法
        nums = [int(x) for x in nums]
        nums.sort()
        return str(nums[-k])
        # 正常做法
        # for i in range(len(nums)):
        #     nums[i] = int(nums[i])
        # nums.sort()
        # return str(nums[-1*k])

    # def quick_sort(self, nums):
    #     left, right = [], []
    #     pivot = nums[0]
    #     for i in range(1, len(nums)):
    #         if nums[i] < pivot:
    #             left.append(nums[i])
    #         else:
    #             right.append(nums[i])
    #     return self.quick_sort(left) + [pivot] + self.quick_sort(right)
# @lc code=end

