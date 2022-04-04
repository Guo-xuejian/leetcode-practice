#
# @lc app=leetcode.cn id=762 lang=python3
#
# [762] 二进制表示中质数个计算置位
#

# @lc code=start
class Solution:
    def countPrimeSetBits(self, left: int, right: int) -> int:
        def isPrime(num):   # 质数判定
            if num < 2:
                return False
            i = 2
            while i * i <= num:
                if num % i == 0:
                    return False
                i += 1
            return True
        def count_bits(num):
            res = 0
            while num:
                if num & 1:
                    res += 1
                num >>= 1
            return res
        res = 0
        for num in range(left, right + 1):
            if isPrime(count_bits(num)):
                res += 1
        return res
# @lc code=end

