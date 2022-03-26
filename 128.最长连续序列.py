#
# @lc app=leetcode.cn id=128 lang=python3
#
# [128] 最长连续序列
#

# @lc code=start
class Solution:
    def longestConsecutive(self, nums: List[int]) -> int:
        num_dict = {}
        res = 0
        for num in nums:
            num_dict[num] = True
        for num in nums:
            if num-1 not in num_dict:
                curr_max = 1
                temp = num + 1
                while temp in num_dict:
                    curr_max += 1
                    temp += 1
                if curr_max > res:
                    res = curr_max
        return res
# @lc code=end

