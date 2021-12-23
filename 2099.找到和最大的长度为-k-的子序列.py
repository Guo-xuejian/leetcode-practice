#
# @lc app=leetcode.cn id=2099 lang=python3
#
# [2099] 找到和最大的长度为 K 的子序列
#

# @lc code=start
class Solution:
    def maxSubsequence(self, nums: List[int], k: int) -> List[int]:
        length = len(nums)
        vals = [[i, nums[i]] for i in range(length)]
        # 按照数值降序排列
        vals.sort(key = lambda x: -x[1])
        vals = sorted(vals[:k])
        # 由于是嵌套列表，所以遍历时加上 idx 并忽略
        res = [val for _, val in vals]
        return res
# @lc code=end

