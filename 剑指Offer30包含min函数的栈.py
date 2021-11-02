class MinStack:

    def __init__(self):
        """
        initialize your data structure here.
        """
        self.stack = []


    def push(self, x: int) -> None:
        self.stack.append(x)


    def pop(self) -> None:
        if self.stack:
            return self.stack.pop()
        return None


    def top(self) -> int:
        a = self.stack.pop()
        if a:
            self.push(a)
        return a


    def min(self) -> int:
        return min(self.stack)




# Your MinStack object will be instantiated and called as such:
# obj = MinStack()
# obj.push(x)
# obj.pop()
# param_3 = obj.top()
# param_4 = obj.min()