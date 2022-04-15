# 剑指 Offer 50. 第一个只出现一次的字符
# 在字符串 s 中找出第一个只出现一次的字符。如果没有，返回一个单空格。 s 只包含小写字母。

# 示例 1:

# 输入：s = "abaccdeff"
# 输出：'b'
# 示例 2:

# 输入：s = "" 
# 输出：' '
 

# 限制：

# 0 <= s 的长度 <= 50000


class Solution:
    def firstUniqChar(self, s: str) -> str:
        if len(s) == 0:
            return " "
        letter_dict = {}
        for letter in s:
            if letter not in letter_dict:
                letter_dict[letter] = 1
            else:
                letter_dict[letter] += 1
        for key, val in letter_dict.items():
            if val == 1:
                return key
        return " "