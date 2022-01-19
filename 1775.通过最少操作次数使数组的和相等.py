#
# @lc app=leetcode.cn id=1775 lang=python3
#
# [1775] 通过最少操作次数使数组的和相等
#

# @lc code=start
class Solution:
    def minOperations(self, nums1: List[int], nums2: List[int]) -> int:
        difference = sum(nums1) - sum(nums2)
        if difference == 0:
            return 0
        # 做一个数字和出现次数的映射
        all_map = [0] * 6
        for num in nums1:
            all_map[num - 1] += 1
        for num in nums2:
            all_map[5 - (num - 1)] += 1
        if difference < 0:
            difference = -difference
            all_map = all_map[::-1]
        res = 0
        for i in range(5, 0, -1):   # 反向，次数最少，就先从反向最大的数字开始
            if all_map[i]*i >= difference:   # 比差值大则只需要修改这个数
                # 向上取整变化次数，也就是5取完了取4321
                return int(res + math.ceil(difference/i))
            difference -= all_map[i]*i
            res += all_map[i]
        return -1
# @lc code=end
