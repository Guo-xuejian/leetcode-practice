#
# @lc app=leetcode.cn id=1001 lang=python3
#
# [1001] 网格照明
#

# @lc code=start
class Solution:
    def gridIllumination(self, n: int, lamps: List[List[int]], queries: List[List[int]]) -> List[int]:
        points = set()
        row, col, diagonal, antiDiagonal = Counter(), Counter(), Counter(), Counter()
        for r, c in lamps:
            if (r, c) in points:
                continue
            points.add((r, c))
            row[r] += 1
            col[c] += 1
            diagonal[r - c] += 1
            antiDiagonal[r + c] += 1

        ans = [0] * len(queries)
        for i, (r, c) in enumerate(queries):
            if row[r] or col[c] or diagonal[r - c] or antiDiagonal[r + c]:
                ans[i] = 1
            for x in range(r - 1, r + 2):
                for y in range(c - 1, c + 2):
                    if x < 0 or y < 0 or x >= n or y >= n or (x, y) not in points:
                        continue
                    points.remove((x, y))
                    row[x] -= 1
                    col[y] -= 1
                    diagonal[x - y] -= 1
                    antiDiagonal[x + y] -= 1
        return ans
# @lc code=end

