class CQueue:

    def __init__(self):
        # 双栈实现队列，stack1为输入栈，stack2为删除栈，stack2在接受删除需求时接收stack1的pop值直至stack1为空，以此维护队列
        self.stack1 = []
        self.stack2 = []

    def appendTail(self, value: int) -> None:
        self.stack1.append(value)

    def deleteHead(self) -> int:
        if self.stack2:
            # 先入先出
            return self.stack2.pop()
        elif self.stack1:
            # 如果删除栈已经为空，那么需要判定输入栈是否还有元素，有的话
            while self.stack1:
                self.stack2.append(self.stack1.pop())
            return self.stack2.pop()
        else:
            return -1



# Your CQueue object will be instantiated and called as such:
# obj = CQueue()
# obj.appendTail(value)
# param_2 = obj.deleteHead()