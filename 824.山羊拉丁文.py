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


# 2022-04-21
# class Solution:
#     def toGoatLatin(self, sentence: str) -> str:
#         vowel_letter = {'a':1,'e':1,'i':1,'o':1,'u':1,'A':1,'E':1,'I':1,'O':1,'U':1}
#         def dfs(word):
#             if word[0] in vowel_letter:
#                 return word + 'ma'
#             else:
#                 return word[1:] + word[0] + 'ma'
#         sentence = sentence.split(" ")
#         for i in range(len(sentence)):
#             sentence[i] = dfs(sentence[i]) + (i + 1) * 'a'
#         return ' '.join(sentence)
# @lc code=end

