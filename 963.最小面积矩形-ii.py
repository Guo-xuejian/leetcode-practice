#
# @lc app=leetcode.cn id=963 lang=python3
#
# [963] 最小面积矩形 II
#

# @lc code=start
class Solution:
    def minAreaFreeRect(self, points: List[List[int]]) -> float:
        n = len(points)
        points_set = set(map(tuple, points))    # 加速查找 hash，空间换取时间
        # 所有的点都不同，每次取3个点，计算第4个点“应该”在哪里，“探”一下存不存在
        res = float('inf')
        for i in range(n):
            [x1, y1] = points[i]
            for j in range(n):
                if i == j:
                    continue
                [x2, y2] = points[j]
                for k in range(j + 1, n):
                    if k == i:
                        continue
                    [x3, y3] = points[k]
                    # p2 p3为副对角线，p1为一个顶角, 矩形的对角线对应点相加和相等
                    x4 = x2 + x3 - x1
                    y4 = y2 + y3 - y1
                    if (x4, y4) in points_set:          # “探”一下，第4个点，是否存在
                        v21 = (x2 - x1,  y2 - y1)
                        v31 = (x3 - x1,  y3 - y1)
                        if v21[0] * v31[0] + v21[1] * v31[1] == 0:  # 判断垂直。2个向量的内积为0
                            cur_area = (v21[0] ** 2 + v21[1] ** 2)**0.5 * \
                                (v31[0] ** 2 + v31[1] ** 2)**0.5
                            if cur_area < res:
                                res = cur_area
        return res if res != float('inf') else 0
# @lc code=end
