#
# @lc app=leetcode.cn id=2034 lang=python3
#
# [2034] 股票价格波动
#

# @lc code=start
from sortedcontainers import SortedList
class StockPrice:
    def __init__(self):
        self.sl = SortedList()  # 有序列表
        self.time_map = {}      # 记录所有时间戳，也可以更新
        self.max_time = -inf    # 最新时间

    def update(self, timestamp: int, price: int) -> None:
        if timestamp in self.time_map:
            # 存在就删除
            self.sl.discard(self.time_map[timestamp])
        self.sl.add(price)
        self.time_map[timestamp] = price
        self.max_time = max(self.max_time, timestamp)

    def current(self) -> int:
        return self.time_map[self.max_time]

    def maximum(self) -> int:
        return self.sl[-1]  # 最大值在最后

    def minimum(self) -> int:
        return self.sl[0]   # 最小值在最前面


# Your StockPrice object will be instantiated and called as such:
# obj = StockPrice()
# obj.update(timestamp,price)
# param_2 = obj.current()
# param_3 = obj.maximum()
# param_4 = obj.minimum()
# @lc code=end

