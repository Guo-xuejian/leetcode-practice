#
# @lc app=leetcode.cn id=335 lang=python3
#
# [335] 路径交叉
#

# @lc code=start
class Solution:
    def isSelfCrossing(self, distance: List[int]) -> bool:
        distance_length = len(distance)
        if distance_length < 4:
            return False
        for i in range(3, distance_length):
            # 第一类交叉===正常的四路交叉 东向移动交叉
            if distance[i] >= distance[i-2] and distance[i-1] < distance[i-3]:
                return True
            
            # 第二类交叉===向北移动正好接上 北向
            if i == 4 and (distance[3] == distance[1] and distance[4] >= distance[2] - distance[0]):
                return True
            
            # 西向交叉 南向交叉
            if i >= 5 and (distance[i-3] - distance[i-5] <= distance[i-1] <= distance[i-3] and distance[i] >= distance[i-2] - distance[i-4] and distance[i-2] > distance[i-4]):
                return True
        return True
# @lc code=end

