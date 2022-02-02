#
# @lc app=leetcode.cn id=2000 lang=python3
#
# [2000] 反转单词前缀
#

# @lc code=start
import os
class Solution:
    def reversePrefix(self, word: str, ch: str) -> str:
        index = -1
        for idx, one in enumerate(word):
            if one == ch:
                index = idx
                break
        
        if index == -1:
            return word
        
        temp, write_index, temp_index = ["" for _ in range(index+1)], 0, index
        while temp_index >= 0:
            temp[write_index] = word[temp_index]
            write_index += 1
            temp_index -= 1
        return "".join(temp) + word[index+1:]
# @lc code=end

