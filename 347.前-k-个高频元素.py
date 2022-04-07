#
# @lc app=leetcode.cn id=347 lang=python3
#
# [347] 前 K 个高频元素
#

# @lc code=start
class Solution:
    def topKFrequent(self, nums: List[int], k: int) -> List[int]:
        if len(nums) == 1:
            return [k]
        v1_dict = {}
        for i in nums:
            v1_dict[i] = v1_dict.get(i, 0) + 1
        data = (sorted(v1_dict.items(), key = lambda x:x[1], reverse=True))[:k]
        return [v[0] for v in data]

# @lc code=end

