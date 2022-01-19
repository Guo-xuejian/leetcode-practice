#
# @lc app=leetcode.cn id=846 lang=python3
#
# [846] 一手顺子
#

# @lc code=start
class Solution:
    def isNStraightHand(self, hand: List[int], groupSize: int) -> bool:
        if len(hand) % groupSize > 0:
            return False
        hand.sort()  # 排序后从小到大
        cnt = Counter(hand)
        for x in hand:
            if cnt[x] == 0:  # 当前数字被抵消，无影响
                continue
            # 对于当前数字后续 groupSize 范围内，匹配每一个应该存在的数字
            for num in range(x, x + groupSize):
                if cnt[num] == 0:   # 无法匹配则 False
                    return False
                cnt[num] -= 1       # 匹配到则抵消一个
        # 所有匹配都完成
        return True
# @lc code=end
