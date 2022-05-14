#
# @lc app=leetcode.cn id=812 lang=python3
#
# [812] 最大三角形面积
#

# @lc code=start
class Solution:
    def getConvexHull(self, points: List[List[int]]) -> List[List[int]]:
        def cross(p: List[int], q: List[int], r: List[int]) -> int:
            return (q[0] - p[0]) * (r[1] - q[1]) - (q[1] - p[1]) * (r[0] - q[0])

        n = len(points)
        if n < 4:
            return points

        # 按照 x 从小到大排序，如果 x 相同，则按照 y 从小到大排序
        points.sort()

        hull = []
        # 求凸包的下半部分
        for i, p in enumerate(points):
            while len(hull) > 1 and cross(hull[-2], hull[-1], p) <= 0:
                hull.pop()
            hull.append(p)
        # 求凸包的上半部分
        m = len(hull)
        for i in range(n - 2, -1, -1):
            while len(hull) > m and cross(hull[-2], hull[-1], points[i]) <= 0:
                hull.pop()
            hull.append(points[i])
        hull.pop()  # hull[0] 同时参与凸包的上半部分检测，因此需去掉重复的 hull[0]
        return hull

    def largestTriangleArea(self, points: List[List[int]]) -> float:
        def triangleArea(x1: int, y1: int, x2: int, y2: int, x3: int, y3: int) -> float:
            return abs(x1 * y2 + x2 * y3 + x3 * y1 - x1 * y3 - x2 * y1 - x3 * y2) / 2

        convexHull = self.getConvexHull(points)
        ans, n = 0, len(convexHull)
        for i, p in enumerate(convexHull):
            k = i + 2
            for j in range(i + 1, n - 1):
                q = convexHull[j]
                while k + 1 < n:
                    curArea = triangleArea(p[0], p[1], q[0], q[1], convexHull[k][0], convexHull[k][1])
                    nextArea = triangleArea(p[0], p[1], q[0], q[1], convexHull[k + 1][0], convexHull[k + 1][1])
                    if curArea >= nextArea:
                        break
                    k += 1
                ans = max(ans, triangleArea(p[0], p[1], q[0], q[1], convexHull[k][0], convexHull[k][1]))
        return ans
# @lc code=end

