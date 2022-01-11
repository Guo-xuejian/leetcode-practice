#
# @lc app=leetcode.cn id=1036 lang=python3
#
# [1036] 逃离大迷宫
#

# @lc code=start
BOUND = int(1e6)


class Solution:
    def isEscapePossible(self, blocked: List[List[int]], source: List[int], target: List[int]) -> bool:
        if len(blocked) == 0:
            return True
        blocked, MAX = {tuple(p) for p in blocked}, len(
            blocked) * (len(blocked) - 1) // 2

        def bfs(start, end):
            points, idx, explored = [start], 0, {tuple(start)}
            while idx < len(points):
                for dx, dy in (0, 1), (1, 0), (-1, 0), (0, -1):  # 下 右 左 上
                    nx, ny = points[idx][0] + \
                        dx, points[idx][1] + dy   # 遍历四个方向
                    # 如果四个方向在限制范围内且没有经过，就加入已探索列表和点位列表
                    if 0 <= nx < BOUND and 0 <= ny < BOUND and (nx, ny) not in blocked and (nx, ny) not in explored:
                        if [nx, ny] == end:
                            return True
                        explored.add((nx, ny))
                        points.append((nx, ny))
                # 点位列表，当包括起点在内的走过的所有点数超过 blpcked 可以围成的最大三角形点数时，也就意味着已经超出了围堵范围，之后一定可以去到 target
                if len(points) > MAX:
                    return True
                idx += 1
            return False

        return bfs(source, target) and bfs(target, source)
# @lc code=end
