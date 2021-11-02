#
# @lc app=leetcode.cn id=453 lang=python3
#
# [453] 最小操作次数使数组元素相等
#

# @lc code=start
class Solution:
    def minMoves(self, nums: List[int]) -> int:
        min_num = min(nums)     # 获取最小值，逐一计算每个数字与最小值的差值加总即可
        res = 0
        for one in nums:
            res += one - min_num
        return res
# @lc code=end

