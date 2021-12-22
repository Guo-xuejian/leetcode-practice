class RecentCounter:

    def __init__(self):
        self.query = []

    def ping(self, t: int) -> int:
        self.query.append(t)
        # 去除不在时间范围内的数据
        while self.query[0] < t - 3000:
            self.query = self.query[1:]
        return len(self.query)



# Your RecentCounter object will be instantiated and called as such:
# obj = RecentCounter()
# param_1 = obj.ping(t)