#
# @lc app=leetcode.cn id=260 lang=python3
#
# [260] 只出现一次的数字 III
#

# @lc code=start
class Solution:
    def singleNumber(self, nums: List[int]) -> List[int]:
        # 排序
        nums = sorted(nums)
        nums_length = len(nums)
        res = []
        temp = nums[0]
        count = 1
        for i in range(1, nums_length):
            if nums[i] == temp:     # 如果和上一元素相等意味着不符合要求，退出
                count += 1
                continue
            # 不相等则有两种情况
            if count < 2:   # 只有 count 小于 2 的情况符合要求，写入结果
                res.append(temp)
            temp = nums[i]  # 覆盖
            count = 1   #计数归 1
        if count < 2:
            res.append(temp)    # 最后一个元素处理
        return res
        
# @lc code=end

