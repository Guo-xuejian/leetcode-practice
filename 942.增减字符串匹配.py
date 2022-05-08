#
# @lc app=leetcode.cn id=942 lang=python3
#
# [942] 增减字符串匹配
#

# @lc code=start
class Solution:
    def diStringMatch(self, S: str) -> List[int]:
        # D 序列递减，I 序列递增
        min_num, max_num = 0, len(S)
        res = []
        for letter in S:
            if letter == 'I':
                # 优先最小，后续递增
                res.append(min_num)
                min_num += 1
            else:
                # 优先最大，后续递减
                res.append(max_num)
                max_num -= 1
        # 最后一个元素没有后续与之比较,加上较小值或者较大值，都可以，最后一定是相等的
        res.append(min_num)
        return res

# 2022-05-09
# class Solution:
#     def diStringMatch(self, s: str) -> List[int]:
#         lo = 0
#         hi = n = len(s)
#         perm = [0] * (n + 1)
#         for i, ch in enumerate(s):
#             if ch == 'I':
#                 perm[i] = lo
#                 lo += 1
#             else:
#                 perm[i] = hi
#                 hi -= 1
#         perm[n] = lo  # 最后剩下一个数，此时 lo == hi
#         return perm

# @lc code=end

