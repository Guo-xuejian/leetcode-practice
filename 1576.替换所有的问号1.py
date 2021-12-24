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
        length = len(s)
        res = ["" for _ in range(length)]
        a_ord = ord("a")
        z_ord = ord("z")
        for i in range(length):
            if s[i] != "?": # 不等于 ? 直接写入
                res[i] = s[i]
                continue
            for j in range(a_ord, z_ord):
                # 边界无前后，其余情况需要判定是否满足与前后元素不相等
                if (i == 0 or chr(j) != res[i-1]) and (i == length - 1 or chr(j) != s[i+1]):
                    res[i] = chr(j) # 转为字符并写入
                    break
        return "".join(res)

# @lc code=end

