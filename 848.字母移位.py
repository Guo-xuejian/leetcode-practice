#
# @lc app=leetcode.cn id=848 lang=python3
#
# [848] 字母移位
#

# @lc code=start
class Solution:
    def shiftingLetters(self, s: str, shifts: List[int]) -> str:
        res = ""
        shifts.append(0)    # 加入一个不产生影响的值，索引处理会更加方便一些
        for i in range(len(s)-1, -1, -1):
            # 循环内叠加数字，单开循环处理会超时
            shifts[i] = (shifts[i] + shifts[i+1]) % 26
            # 对应字符移位写入 res, 需要注意是个循环，所以需要取余
            res = chr(ord("a") + (ord(s[i]) + shifts[i] - ord("a")) % 26) + res
        return res
# @lc code=end
