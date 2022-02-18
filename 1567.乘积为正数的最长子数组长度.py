#
# @lc app=leetcode.cn id=1567 lang=python3
#
# [1567] 乘积为正数的最长子数组长度
#

# @lc code=start
class Solution:
    def getMaxLen(self, nums: List[int]) -> int:
        m = len(nums)
        positive, negative = [0] * m, [0] * m
        if nums[0] > 0:
            positive[0] = 1
        elif nums[0] < 0:
            negative[0] = 1

        maxLength = positive[0]
        for i in range(1, m):
            if nums[i] > 0:
                positive[i] = positive[i - 1] + 1   # 大于正数长度加一
                negative[i] = (negative[i - 1] +
                               1 if negative[i - 1] > 0 else 0)   # 负数同样
            elif nums[i] < 0:
                # 小于 0 时，始终保持 positive 为正数长度，切换 negative 至 positive
                positive[i] = (negative[i - 1] +
                               1 if negative[i - 1] > 0 else 0)   # 未出现负数置为 0
                negative[i] = positive[i - 1] + 1
            else:   # 0 则切分
                positive[i] = negative[i] = 0
            maxLength = max(maxLength, positive[i])  # 始终以 positive 做比较

        return maxLength

# @lc code=end
