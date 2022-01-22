#
# @lc app=leetcode.cn id=1332 lang=python3
#
# [1332] 删除回文子序列
#

# @lc code=start
class Solution:
    def removePalindromeSub(self, s: str) -> int:
        # 最多也就两次，由于一种字符组成的字符串为回文，所以存在两种字符时直接删除其中一种，那么剩下的另一种自然就是回文
        # 我们只需要判定原字符串是不是回文即可
        if s == s[::-1]:
            return 1
        else:
            return 2
# @lc code=end

