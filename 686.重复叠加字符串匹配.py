#
# @lc app=leetcode.cn id=686 lang=python3
#
# [686] 重复叠加字符串匹配
#

# @lc code=start
class Solution:
    def repeatedStringMatch(self, a: str, b: str) -> int:
        res = 1
        if b in a:
            return res
        for one in b:
            if one not in a:
                return -1
        # 所谓三倍都没有的话那就是无了
        max_length = 2*len(a) + len(b)
        temp_string = ""
        while len(temp_string) <= max_length:
            res += 1
            temp_string = a*res
            if b in temp_string:
                return res
        return -1
# @lc code=end

