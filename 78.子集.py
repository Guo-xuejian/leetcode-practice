#
# @lc app=leetcode.cn id=78 lang=python3
#
# [78] 子集
#

# @lc code=start
class Solution:
    def subsets(self, nums: List[int]) -> List[List[int]]:
        res = [[], ]
        for num in nums:
            curr = []
            for part in res:
                curr.append(part[:] + [num, ])
            res += curr
        return res
# @lc code=end

