#
# @lc app=leetcode.cn id=315 lang=python3
#
# [315] 计算右侧小于当前元素的个数
#

# @lc code=start
class Solution:
    def countSmaller(self, nums: List[int]) -> List[int]:
        def getRightSmallNum(start, nums, length):  # 获取右边小于该值的数量
            smaller_count = 0   # 初始化为 0，即使处于边界也返回正确值
            starter = nums[start]
            for i in range(start, length):
                if nums[i] < starter:
                    smaller_count += 1
            return smaller_count

        nums_length = len(nums)
        if nums_length == 1:    # 长度为 0 时不会有右边的值，直接返回
            return [0]
        # result_list = [0 for i in range(nums_length)]
        for i in range(nums_length):
            # 直接在原数组上进行修改，降低空间复杂度
            nums[i] = getRightSmallNum(i, nums, nums_length)
        return nums
# @lc code=end
