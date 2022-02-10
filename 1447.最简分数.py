#
# @lc app=leetcode.cn id=1447 lang=python3
#
# [1447] 最简分数
#

# @lc code=start
class Solution:
    def simplifiedFractions(self, n: int) -> List[str]:
        res = []
        for down in range(2, n + 1):
            for up in range(down):
                if self.gcd(up, down) == 1:
                    res.append(f"{up}/{down}")
        return res
    
    def gcd(self, x, y):
        while x != 0:
            x, y = y % x, x
        return y
# @lc code=end

