#
# @lc app=leetcode.cn id=475 lang=python3
#
# [475] 供暖器
#

# @lc code=start
class Solution:
    def findRadius(self, houses: List[int], heaters: List[int]) -> int:
        heaters.sort()
        res = 0
        for house in houses:
            j = bisect_right(heaters, house)    # 靠右选择
            i = j - 1
            rightDistance = heaters[j] - house if j < len(heaters) else float('inf')
            leftDistance = house - heaters[i] if i >= 0 else float('inf')
            # 比较取较小
            curDistance = min(leftDistance, rightDistance)
            # 最后取较大
            res = max(res, curDistance)
        return res
# @lc code=end

