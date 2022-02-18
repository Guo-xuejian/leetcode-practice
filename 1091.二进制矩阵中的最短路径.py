#
# @lc app=leetcode.cn id=1091 lang=python3
#
# [1091] 二进制矩阵中的最短路径
#

# @lc code=start
from collections import deque
class Solution:
    def shortestPathBinaryMatrix(self, grid: List[List[int]]) -> int:
        if grid[0][0] != 0:
            return -1
        length = len(grid)
        if length == 1:
            return 1
        que = deque()
        visited = {}   # 已访问的路径
        que.appendleft((0,0))
        visited[(0,0)] = True
        start = 1
        while que:
            for _ in range(len(que)):
                ind, con = que.pop()
                # 八个方向
                for pos_h, pos_v in [(-1,-1), (-1,0), (-1,1), (0,1), (1,1), (1,0), (1,-1), (0,-1)]:
                    new_ind = ind + pos_h
                    new_con = con + pos_v
                    # 并未越界同时为 0 ，且未访问
                    if 0 <= new_ind < length and 0 <= new_con < length and grid[new_ind][new_con] == 0 and not visited.get((new_ind, new_con)):
                        # 已到终点，+1 返回
                        if new_ind == length - 1 and new_con == length - 1:
                            return start + 1
                        que.appendleft((new_ind, new_con))
                        visited[(new_ind, new_con)] = True  # 已访问
            start += 1  # 正常 +1
        return -1
# @lc code=end

