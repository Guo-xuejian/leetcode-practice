#
# @lc app=leetcode.cn id=2085 lang=python3
#
# [2085] 统计出现过一次的公共字符串
#

# @lc code=start
class Solution:
    def countWords(self, words1: List[str], words2: List[str]) -> int:
        res = 0
        words1_dict = {}
        words2_dict = {}
        for one in words1:
            if one in words1_dict:
                words1_dict[one] += 1
            else:
                words1_dict[one] = 1
        for one in words2:
            if one in words2_dict:
                words2_dict[one] += 1
            else:
                words2_dict[one] = 1
        for key, val in words1_dict.items():
            if val == 1 and key in words2_dict and words2_dict[key] == 1:
                res += 1
        return res
# @lc code=end

