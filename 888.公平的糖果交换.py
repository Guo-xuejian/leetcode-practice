#
# @lc app=leetcode.cn id=888 lang=python3
#
# [888] 公平的糖果交换
#

# @lc code=start
class Solution:
    def fairCandySwap(self, aliceSizes: List[int], bobSizes: List[int]) -> List[int]:
        sum_alice, sum_bob = sum(aliceSizes), sum(bobSizes)
        delta = (sum_alice - sum_bob) // 2
        rec = set(aliceSizes)
        res = None
        # 差值，也就是较小一方中整体差 delta
        # 只要有一个数满足加上delta在较大方找到
        # 那就直接交换对应数字即可，就可以满足双方相等
        for y in bobSizes:
            x = y + delta
            if x in rec:
                res = [x, y]
                break
        return res
# @lc code=end

