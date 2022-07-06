#
# @lc app=leetcode.cn id=648 lang=python3
#
# [648] 单词替换
#

# @lc code=start
class Solution:
    def replaceWords(self, dictionary: List[str], sentence: str) -> str:
        dictionarySet = set(dictionary)
        words = sentence.split(' ')
        for i, word in enumerate(words):
            for j in range(1, len(words) + 1):
                if word[:j] in dictionarySet:
                    words[i] = word[:j]
                    break
        return ' '.join(words)
# @lc code=end

