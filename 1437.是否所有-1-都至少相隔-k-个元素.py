#
# @lc app=leetcode.cn id=1437 lang=python3
#
# [1437] 是否所有 1 都至少相隔 k 个元素
#

# @lc code=start
class Solution:
    def kLengthApart(self, nums: List[int], k: int) -> bool:
        if k == 0:  # 间隔 0 天生满足
            return True
        if len(nums) < k:   # 等于不行，如果 100，只有一个 1，那么也是符合要求的
            return False
        pre_index = 0
        exist_status = False
        for idx, val in enumerate(nums):
            if val == 1:
                if not exist_status:    # 还没出现 1
                    exist_status = True
                    pre_index = idx
                    continue
                if idx - pre_index - 1 >= k:
                    pre_index = idx
                else:
                    return False
        return True
# @lc code=end
