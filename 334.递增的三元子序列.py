#
# @lc app=leetcode.cn id=334 lang=python3
#
# [334] 递增的三元子序列
#

# @lc code=start
class Solution:
    def increasingTriplet(self, nums: List[int]) -> bool:
        length = len(nums)
        if length < 3:
            return False
        min_list = [0 for _ in range(length)]
        max_list = [0 for _ in range(length)]
        i, j = 0, length-1
        while i < length and j >= 0:
            if i == 0:
                min_list[i] = nums[0]
            else:
                min_list[i] = min(nums[i], min_list[i-1])
            i += 1
            if j == length - 1:
                max_list[j] = nums[j]
            else:
                max_list[j] = max(nums[j], max_list[j+1])
            j -= 1
        print(min_list)
        print(max_list)
        for i in range(1, length):
            if nums[i] > min_list[i] and nums[i] < max_list[i]:
                return True
        return False
# @lc code=end
