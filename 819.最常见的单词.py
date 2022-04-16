#
# @lc app=leetcode.cn id=819 lang=python3
#
# [819] 最常见的单词
#

# @lc code=start
class Solution:
    def mostCommonWord(self, paragraph: str, banned: List[str]) -> str:
        banned_dict = {}
        for word in banned:
            banned_dict[word] = 1
        paragraph_dict = {}
        index, pre = 0, 0
        while index < len(paragraph):
            if not paragraph[index].isalpha():
                curr = paragraph[pre:index].lower()
                if curr not in banned_dict:
                    if curr not in paragraph_dict:
                        paragraph_dict[curr] = 1
                    else:
                        paragraph_dict[curr] += 1
                pre = index + 1
            index += 1
        if pre != len(paragraph):
            curr = paragraph[pre:].lower()
            if curr not in banned_dict and curr not in paragraph_dict:
                paragraph_dict[curr] = 1
            else:
                paragraph_dict[curr] += 1
        max_times = 0
        res = ""
        for key, val in paragraph_dict.items():
            if key != "" and val > max_times:
                res = key
                max_times = val
        return res
# @lc code=end

