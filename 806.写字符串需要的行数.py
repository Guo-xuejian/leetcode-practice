#
# @lc app=leetcode.cn id=806 lang=python3
#
# [806] 写字符串需要的行数
#

# @lc code=start
class Solution:
    def numberOfLines(self, widths: List[int], s: str) -> List[int]:
        curr_length, lines = 0, 0
        ord_a = ord("a")
        for letter in s:
            curr = widths[ord(letter) - ord_a]
            if curr_length + curr > 100:
                lines += 1
                curr_length = curr
            else:
                curr_length += curr
        return [lines + 1, curr_length]
# @lc code=end

