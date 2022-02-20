# 剑指 Offer 13. 机器人的运动范围
# 地上有一个m行n列的方格，从坐标 [0,0] 到坐标 [m-1,n-1] 。一个机器人从坐标 [0, 0] 的格子开始移动，它每次可以向左、右、上、下移动一格（不能移动到方格外），也不能进入行坐标和列坐标的数位之和大于k的格子。例如，当k为18时，机器人能够进入方格 [35, 37] ，因为3+5+3+7=18。但它不能进入方格 [35, 38]，因为3+5+3+8=19。请问该机器人能够到达多少个格子？

#  

# 示例 1：

# 输入：m = 2, n = 3, k = 1
# 输出：3
# 示例 2：

# 输入：m = 3, n = 1, k = 0
# 输出：1
# 提示：

# 1 <= n,m <= 100
# 0 <= k <= 20


class Solution:
    def movingCount(self, m: int, n: int, k: int) -> int:
        def is_valid(x, y):
            if not (0 <= x < m and 0 <= y < n): # 不可越界
                return False
            judge_sum = 0
            while x:
                judge_sum += x % 10
                x = x // 10
            while y:
                judge_sum += y % 10
                y = y // 10
            if judge_sum > k:   # 数位和不能大于 k
                return False
            return True
        # bfs
        q = [(0, 0)]
        visited = {(0, 0):True}
        directions = [(1, 0), (-1, 0), (0, 1), (0, -1)]
        res = 1
        while q:
            length = len(q)
            for _ in range(length):
                point = q[0]
                q = q[1:]
                for direction in directions:
                    x, y = point[0] + direction[0], point[1] + direction[1]
                    if is_valid(x, y) and (x, y) not in visited:
                        res += 1
                        visited[(x, y)] = True
                        q.append((x, y))
        return res