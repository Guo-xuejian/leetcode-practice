#
# @lc app=leetcode.cn id=709 lang=python3
#
# [709] 转换成小写字母
#

# @lc code=start
class Solution:
    def toLowerCase(self, s: str) -> str:
        # return s.lower()
        res = ["" for _ in range(len(s))]
        A_ord, Z_ord = ord("A"), ord("Z")
        for letter in s:
            curr_ord = ord(letter)
            if A_ord <= curr_ord <= Z_ord:
                res.append(chr(curr_ord ^ (32)))
            else:
                res.append(letter)
        return "".join(res)
# @lc code=end

