#
# @lc app=leetcode.cn id=733 lang=python3
#
# [733] 图像渲染
#

# @lc code=start
class Solution:
    def floodFill(self, image: List[List[int]], sr: int, sc: int, newColor: int) -> List[List[int]]:
        m, n = len(image), len(image[0])
        preColor = image[sr][sc]
        # 四个方向
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]

        def dfs(x, y):
            # 索引不能越界
            if x < 0 or y < 0 or x >= m or y >= n:
                return
            if image[x][y] == preColor:     # 和初始坐标原色相同就修改
                image[x][y] = newColor
                for direction in directions:    # 遍历四个方向
                    dfs(x + direction[0], y + direction[1])

        if newColor != preColor:
            dfs(sr, sc)
        return image
# @lc code=end
