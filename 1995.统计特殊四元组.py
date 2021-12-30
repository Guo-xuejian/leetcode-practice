#
# @lc app=leetcode.cn id=1995 lang=python3
#
# [1995] 统计特殊四元组
#

# @lc code=start
class Solution:
    def countQuadruplets(self, nums: List[int]) -> int:
        res = 0
        nums_length = len(nums)
        for a in range(nums_length):
            for b in range(a+1, nums_length):
                for c in range(b+1, nums_length):
                    for d in range(c+1, nums_length):
                        if nums[a] + nums[b] + nums[c] == nums[d]:
                            res += 1
        return res
# @lc code=end

