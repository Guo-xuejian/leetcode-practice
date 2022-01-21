#
# @lc app=leetcode.cn id=2029 lang=python3
#
# [2029] 石子游戏 IX
#

# @lc code=start
class Solution:
    def stoneGameIX(self, stones: List[int]) -> bool:
        cnt0 = cnt1 = cnt2 = 0
        for val in stones:
            if (typ := val % 3) == 0:   # 倍数
                cnt0 += 1
            elif typ == 1:              # 余 1
                cnt1 += 1
            else:                       # 余 2
                cnt2 += 1
        if cnt0 % 2 == 0:
            return cnt1 >= 1 and cnt2 >= 1
        return cnt1 - cnt2 > 2 or cnt2 - cnt1 > 2 
# @lc code=end

