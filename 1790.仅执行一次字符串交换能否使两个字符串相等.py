#
# @lc app=leetcode.cn id=1790 lang=python3
#
# [1790] 仅执行一次字符串交换能否使两个字符串相等
#

# @lc code=start
class Solution:
    def areAlmostEqual(self, s1: str, s2: str) -> bool:
        res = []
        for i in range(len(s1)):
            if s1[i] != s2[i]:  # 写入不同字符索引
                res.append(i)
            if len(res) > 2:    # 长度超过 2 时退出，必定不满足
                return False
        if len(res) == 1:   # 1 只有一处不相同，不满足
            return False
        # 无不同 或者 两处不同且互补
        if len(res) == 0 or (s1[res[0]] == s2[res[1]] and s1[res[1]] == s2[res[0]]):
            return True
        return False
# @lc code=end
