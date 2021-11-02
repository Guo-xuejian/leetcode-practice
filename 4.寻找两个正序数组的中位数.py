#
# @lc app=leetcode.cn id=4 lang=python3
#
# [4] 寻找两个正序数组的中位数
#

# @lc code=start
class Solution:
    def findMedianSortedArrays(self, nums1: List[int], nums2: List[int]) -> float:
        # 不需要去重
        order_group_nums = sorted(nums1 + nums2)
        mid_num = len(order_group_nums) / 2
        if int(mid_num) - mid_num == 0:
            # 说明数组元素总数量为偶数，中位数为 /2
            return (order_group_nums[int(mid_num) - 1] + order_group_nums[int(mid_num)]) / 2
        else:
            # 说明为奇数，中位数向下取整
            return order_group_nums[int(mid_num)]
        
# @lc code=end

