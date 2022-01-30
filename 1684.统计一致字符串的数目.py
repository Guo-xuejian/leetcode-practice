#
# @lc app=leetcode.cn id=1684 lang=python3
#
# [1684] 统计一致字符串的数目
#

# @lc code=start
class Solution:
    def countConsistentStrings(self, allowed: str, words: List[str]) -> int:
        res = 0
        allowed_letter_dict = {}
        for letter in allowed:
            allowed_letter_dict[letter] = 1
        
        for word in words:
            exist_status = False
            for letter in word:
                if letter not in allowed_letter_dict:
                    exist_status = True
                    break
            if not exist_status:
                res += 1
        return res
# @lc code=end

