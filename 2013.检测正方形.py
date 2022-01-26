#
# @lc app=leetcode.cn id=2013 lang=python3
#
# [2013] 检测正方形
#

# @lc code=start
class DetectSquares:

    def __init__(self):
        self.map = defaultdict(Counter)

    def add(self, point: List[int]) -> None:
        x, y = point
        self.map[y][x] +=  1


    def count(self, point: List[int]) -> int:
        res = 0
        x,y = point
        if not y in self.map:
            return 0
        yCnt = self.map[y]      # 当前 y 坐标轴上的所有点 x 的字典

        for col, colCnt in self.map.items():
            if col != y:        # 与当前 y 坐标不同
                d = col - y     # 宽 === 正方形边长
                # colCnt[x] 当前点数量   yCnt[x+d] 右方点数量  colCnt[x+d]目前点上方点数量
                res += colCnt[x] * yCnt[x+d] * colCnt[x+d]
                res += colCnt[x] * yCnt[x-d] * colCnt[x-d]


# Your DetectSquares object will be instantiated and called as such:
# obj = DetectSquares()
# obj.add(point)
# param_2 = obj.count(point)
# @lc code=end

