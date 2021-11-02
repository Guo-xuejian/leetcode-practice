#
# @lc app=leetcode.cn id=229 lang=python3
#
# [229] 求众数 II
#

# @lc code=start
class Solution:
    def majorityElement(self, nums: List[int]) -> List[int]:
        nums_length = len(nums)
        three_limit_num = int(nums_length/3)
        nums = sorted(nums)
        temp = nums[0]
        count = 1
        res = []
        for i in range(1, nums_length):     # 1 剔除了第一个元素多计数一次
            if nums[i] == temp:     # 和之前的元素相等
                count += 1      # 计数加一
                continue
            if count > three_limit_num:      # 不相等且满足条件
                res.append(temp)    # 加入结果列表
            # 覆盖 temp 为新值，将 count 归 1
            temp = nums[i]
            count = 1
        # 数组的最后几个都为同一元素会导致条件判定未发生，需要特殊处理
        if count > three_limit_num:
            res.append(temp)
        return res
# @lc code=end

