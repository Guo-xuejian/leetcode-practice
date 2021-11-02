#
# @lc app=leetcode.cn id=187 lang=python3
#
# [187] 重复的DNA序列
#

# @lc code=start
class Solution:
    def findRepeatedDnaSequences(self, s: str) -> List[str]:
        string_dict = {}
        for i in range(len(s)-10 +1):
            if len(s)+1-i >= 10:
                if s[i:i+10] in string_dict:
                    string_dict[s[i:i+10]] += 1
                else:
                    string_dict[s[i:i+10]] = 1
        final_res = []
        for i,val in string_dict.items():
            if val > 1:
                final_res.append(i)
        return final_res
# @lc code=end

