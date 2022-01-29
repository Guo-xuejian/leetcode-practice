#
# @lc app=leetcode.cn id=824 lang=python3
#
# [824] 山羊拉丁文
#

# @lc code=start

alphabet = ["a", "e", "i", "o", "u", "A", "E", "I", "O", "U"]

class Solution:
    def toGoatLatin(self, sentence: str) -> str:
        sentence = [self.transfer(idx + 1, word) for idx, word in enumerate(sentence.split())]
        return " ".join(sentence)
    
    def transfer(self, index, word):
        if word[0] not in alphabet:
            word = word[1:] + word[0]
        word += "ma" + "a" * index
        return word
# @lc code=end

