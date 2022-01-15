# 剑指 Offer II 004. 只出现一次的数字 
# 给你一个整数数组 nums ，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次 。请你找出并返回那个只出现了一次的元素。

# 示例 1：
# 输入：nums = [2,2,3,2]
# 输出：3

# 示例 2：
# 输入：nums = [0,1,0,1,0,1,100]
# 输出：100
 
# 提示：
# 1 <= nums.length <= 3 * 104
# -231 <= nums[i] <= 231 - 1
# nums 中，除某个元素仅出现 一次 外，其余每个元素都恰出现 三次

class Solution:
    def singleNumber(self, nums: List[int]) -> int:
        # 字典实现
        # 执行用时：48 ms, 在所有 Python3 提交中击败了73.16%的用户
        # 内存消耗：15.9 MB, 在所有 Python3 提交中击败了28.04%的用户
        num_times_dict = {}
        for num in nums:
            if num in num_times_dict:
                num_times_dict[num] += 1
            else:
                num_times_dict[num] = 1
        for key, val in num_times_dict.items():
            if val == 1:
                return key

        # 排序实现
        # 执行用时：48 ms, 在所有 Python3 提交中击败了38.46%的用户
        # 内存消耗：15.9 MB, 在所有 Python3 提交中击败了58.33%的用户
        # nums.sort()
        # pre_times = 1
        # pre_num = nums[0]
        # for i in range(1, len(nums)):
        #     if nums[i] != nums[i-1]:
        #         if pre_times == 1:  # 前一个数字字数满足要求
        #             return pre_num
        #         else:   # 不满足重置
        #             pre_times = 1
        #             pre_num = nums[i]
        #     else:
        #         pre_times += 1
        # # 走到这里实际上就是最后一个，但是没法比较，直接返回
        # return pre_num