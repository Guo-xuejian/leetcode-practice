#
# @lc app=leetcode.cn id=49 lang=python3
#
# [49] 字母异位词分组
#

# @lc code=start
class Solution:
    def groupAnagrams(self, strs: List[str]) -> List[List[str]]:
        all_res = {}
        for word in strs:
            key = "".join(sorted(list(word)))
            if key in all_res:
                all_res[key].append(word)
            else:
                all_res[key] = [word, ]
        return [val for _, val in all_res.items()]
# @lc code=end

