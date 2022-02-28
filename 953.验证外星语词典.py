#
# @lc app=leetcode.cn id=953 lang=python3
#
# [953] 验证外星语词典
#

# @lc code=start
class Solution:
    def isAlienSorted(self, words: List[str], order: str) -> bool:
        letter_index = {order[i]: i for i in range(len(order))}

        for i in range(len(words) - 1):
            word1, word2 = words[i], words[i + 1]
            is_prefix = True
            for j in range(min(len(word1), len(word2))):
                if word1[j] != word2[j]:
                    if letter_index[word1[j]] > letter_index[word2[j]]:
                        return False
                    is_prefix = False
                    break
            if is_prefix and len(word1) > len(word2):
                return False
        return True
# @lc code=end

