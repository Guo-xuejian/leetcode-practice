#
# @lc app=leetcode.cn id=1281 lang=python3
#
# [1281] 整数的各位积和之差
#

# @lc code=start
class Solution:
    def subtractProductAndSum(self, n: int) -> int:
        multi_num, sum_num = 1, 0
        while n:
            curr = n % 10
            multi_num *= curr
            sum_num += curr
            n = n // 10
        return multi_num - sum_num
# @lc code=end

