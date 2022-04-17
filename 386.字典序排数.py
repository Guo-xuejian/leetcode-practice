#
# @lc app=leetcode.cn id=386 lang=python3
#
# [386] 字典序排数
#

# @lc code=start
class Solution:
    def lexicalOrder(self, n: int) -> List[int]:
        res = [0 for _ in range(n)]
        num = 1
        for i in range(n):
            res[i] = num
            if num * 10 <= n:
                num *= 10
            else:
                while num % 10 == 9 or num + 1 > n:
                    num //= 10
                num += 1
        return res
# @lc code=end

