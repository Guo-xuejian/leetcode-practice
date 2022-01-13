#
# @lc app=leetcode.cn id=2124 lang=python3
#
# [2124] 检查是否所有 A 都在 B 之前
#

# @lc code=start
class Solution:
    def checkString(self, s: str) -> bool:
        a_status = False
        b_status = False
        for i in range(len(s)):
            if s[i] == "a":
                if b_status:
                    return False
                if not a_status:
                    a_status = True
            else:
                if not b_status:
                    b_status = True
        return True

# @lc code=end
