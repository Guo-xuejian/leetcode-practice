#
# @lc app=leetcode.cn id=392 lang=python3
#
# [392] 判断子序列
#

# @lc code=start
class Solution:
    def isSubsequence(self, s: str, t: str) -> bool:
        if not s:
            return True
        elif not t and s:
            return False
        # 维护当前字母和上一个字母的索引
        pre_index = -1
        after_index = -1

        s_length = len(s)
        t_length = len(t)
        for i in range(s_length):
            exist_status = False
            for j in range(t_length):
                if s[i] == t[j]:
                    exist_status = True
                    if j > after_index: # 只要出现一个满足条件的即可
                        after_index = j
                        break
            # 存在且索引符合条件则重置索引，
            # 否则直接返回 false 因为不存在或者索引不满足条件
            if exist_status and after_index > pre_index:
                pre_index = after_index
            else:
                return False
        return True
# @lc code=end

