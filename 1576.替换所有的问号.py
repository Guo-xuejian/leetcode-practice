#
# @lc app=leetcode.cn id=1576 lang=python3
#
# [1576] 替换所有的问号
#

# @lc code=start
class Solution:
    def modifyString(self, s):
        if s == '?':
            return 'a'

        def comp(x, y='A'):
            ox, oy = ord(x), ord(y)
            for num in range(97, 123):
                if num != ox and num != oy:
                    return chr(num)

        ret = ""
        for i, j in enumerate(s):
            if j == '?':
                if i == 0:
                    j = comp(s[1])
                elif i == len(s) - 1:
                    j = comp(ret[-1])
                else:
                    j = comp(ret[-1], s[i + 1])
            ret += j
        return ret

# @lc code=end

