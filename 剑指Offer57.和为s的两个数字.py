# 剑指 Offer 57. 和为s的两个数字
# 输入一个递增排序的数组和一个数字s，在数组中查找两个数，使得它们的和正好是s。如果有多对数字的和等于s，则输出任意一对即可。

# 示例 1：
# 输入：nums = [2,7,11,15], target = 9
# 输出：[2,7] 或者 [7,2]
# 示例 2：

# 输入：nums = [10,26,30,31,47,60], target = 40
# 输出：[10,30] 或者 [30,10]
# 限制：

# 1 <= nums.length <= 10^5
# 1 <= nums[i] <= 10^6

class Solution:
    def twoSum(self, nums: List[int], target: int) -> List[int]:
        length = len(nums)
        if length == 1:
            return []
        # 双指针
        left, right = 0, length - 1
        while left < right:
            sum = nums[left] + nums[right]
            if sum > target:    # 大于让较大的 right 左移一位
                right -= 1
            elif sum < target:  # 小于让较小的 left右移一位
                left += 1
            else:   # 相等返回
                return [nums[left], nums[right]]
        # for i in range(length-1):     # 超时
        #     for j in range(i+1, length):
        #         if nums[i] + nums[j] == target:
        #             return [nums[i], nums[j]]
        return []
    # def quick_sort(self, nums):
    #     left, right = [], []
    #     pivot = nums[0]
    #     for i in range(1,len(nums)):
    #         if nums[i] < pivot:
    #             left.append(nums[i])
    #         if nums[i] >= pivot:
    #             right.append(nums[i])
    #     return self.quick_sort(left) + [pivot] + self.quick_sort(right)
