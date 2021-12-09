#
# @lc app=leetcode.cn id=397 lang=python3
#
# [397] 整数替换
#

# @lc code=start
class Solution:
    def integerReplacement(self, n: int) -> int:
        if n == 1:  # 1 满足条件不再处理
            return 0
        # 偶数加 1（除法的一步）递归
        if n % 2 == 0:  # 必须是 // 不然python会出现小数
            return 1 + self.integerReplacement(n//2)
        # 奇数之后必定为偶，则直接加 2（加法之后做除法共计两步）继续递归
        return 2 + min(self.integerReplacement(n//2), self.integerReplacement(n//2+1))
# @lc code=end
