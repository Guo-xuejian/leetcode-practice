#
# @lc app=leetcode.cn id=140 lang=python3
#
# [140] 单词拆分 II
#

# @lc code=start
class Solution:
    def wordBreak(self, s: str, wordDict: List[str]) -> List[str]:
        word_dict = {}
        for word in wordDict:
            word_dict[word] = True
        
        res = []
        def check(s, temp_list):
            if not s:   # 为空，完全匹配成功，拼接后写入结果集
                res.append(" ".join(temp_list))
                return
            for i in range(len(s) + 1):
                if s[:i] in word_dict:  # 当前字符串存在，记录并递归之后的字符串
                    check(s[i:], temp_list + [s[:i]])
        
        check(s, [])
        return res
# @lc code=end

