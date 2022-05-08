#
# @lc app=leetcode.cn id=442 lang=python3
#
# [442] 数组中重复的数据
#

# @lc code=start
class Solution:
    def findDuplicates(self, nums: List[int]) -> List[int]:
        # 字典
        num_time_dict = {}
        res = []
        for num in nums:
            if num not in num_time_dict:
                num_time_dict[num] = 1
            else:
                res.append(num)
        return res
# @lc code=end

