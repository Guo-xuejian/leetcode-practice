# 面试题 01.02. 判定是否互为字符重排
# 给定两个字符串 s1 和 s2，请编写一个程序，确定其中一个字符串的字符重新排列后，能否变成另一个字符串。

# 示例 1：

# 输入: s1 = "abc", s2 = "bca"
# 输出: true 
# 示例 2：

# 输入: s1 = "abc", s2 = "bad"
# 输出: false
# 说明：

# 0 <= len(s1) <= 100
# 0 <= len(s2) <= 100

class Solution:
    def CheckPermutation(self, s1: str, s2: str) -> bool:
        if len(s1) != len(s2):
            return False
        letter_dict = {}
        # 两方抵消，最后不为 0 ， 则False
        for letter in s1:
            if letter in letter_dict:
                letter_dict[letter] += 1
            else:
                letter_dict[letter] = 1
        for letter in s2:
            if letter in letter_dict:
                letter_dict[letter] -= 1
            else:
                letter_dict[letter] = 1
        for value in letter_dict.values():
            if value != 0:
                return False
        return True