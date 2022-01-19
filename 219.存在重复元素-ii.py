#
# @lc app=leetcode.cn id=219 lang=python3
#
# [219] 存在重复元素 II
#

# @lc code=start
class Solution:
    def containsNearbyDuplicate(self, nums: List[int], k: int) -> bool:
        # 以空间换时间
        all_dict = {}
        for i, v in enumerate(nums):
            # 从前往后，直接相减就是 abs
            if v in all_dict and i - all_dict[v] <= k:
                return True
            all_dict[v] = i
        return False
# @lc code=end

