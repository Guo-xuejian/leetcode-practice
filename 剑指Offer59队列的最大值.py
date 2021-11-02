class MaxQueue:

    def __init__(self):
        self.queue = []


    def max_value(self) -> int:
        if self.queue:
            return max(self.queue)
        return -1


    def push_back(self, value: int) -> None:
        self.queue.insert(0, value)


    def pop_front(self) -> int:
        if self.queue:
            return self.queue.pop()
        return -1


# Your MaxQueue object will be instantiated and called as such:
# obj = MaxQueue()
# param_1 = obj.max_value()
# obj.push_back(value)
# param_3 = obj.pop_front()