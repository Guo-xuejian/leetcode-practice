# 剑指 Offer 09. 用两个栈实现队列
# 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

# 示例 1：

# 输入：
# ["CQueue","appendTail","deleteHead","deleteHead"]
# [[],[3],[],[]]
# 输出：[null,null,3,-1]
# 示例 2：

# 输入：
# ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
# [[],[],[5],[2],[],[]]
# 输出：[null,-1,null,null,5,2]
# 提示：

# 1 <= values <= 10000
# 最多会对 appendTail、deleteHead 进行 10000 次调用

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