#
# @lc app=leetcode.cn id=338 lang=python3
#
# [338] 比特位计数
#

# @lc code=start
class Solution:
    def countBits(self, n: int) -> List[int]:
        def countOnes(x):
            ones = 0
            while x > 0:
                x &= x - 1
                ones += 1
            return ones
        
        bits = [countOnes(i) for i in range(n+1)]
        return bits
# @lc code=end

