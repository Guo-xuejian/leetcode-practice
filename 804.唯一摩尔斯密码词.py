#
# @lc app=leetcode.cn id=804 lang=python3
#
# [804] 唯一摩尔斯密码词
#

# @lc code=start
class Solution:
    def uniqueMorseRepresentations(self, words: List[str]) -> int:
        morse_code = [".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--","-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--.."]
        code_dict = {}
        a_code = ord("a")
        for i in range(26):
            code_dict[chr(i + a_code)] = morse_code[i]
        res = 0
        res_dict = {}
        for word in words:
            curr = ""
            for letter in word:
                curr += code_dict[letter]
            if curr not in code_dict:
                res += 1
                code_dict[curr] = 1
        return res
# @lc code=end

