#
# @lc app=leetcode.cn id=1716 lang=python3
#
# [1716] 计算力扣银行的钱
#

# @lc code=start
class Solution:
    def totalMoney(self, n: int) -> int:
        res = 0
        currDay = 0 # 虚拟一个第零天，后续相加
        for i in range(n):
            currDay += 1
            if i % 7 == 0 and i != 0:
                currDay -= 6    # -7+1 ===> -6
            res += currDay
        return res
# @lc code=end

