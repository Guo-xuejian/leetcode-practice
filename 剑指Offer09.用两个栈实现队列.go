// 剑指 Offer 09. 用两个栈实现队列
// 用两个栈实现一个队列。队列的声明如下，请实现它的两个函数 appendTail 和 deleteHead ，分别完成在队列尾部插入整数和在队列头部删除整数的功能。(若队列中没有元素，deleteHead 操作返回 -1 )

 

// 示例 1：

// 输入：
// ["CQueue","appendTail","deleteHead","deleteHead"]
// [[],[3],[],[]]
// 输出：[null,null,3,-1]
// 示例 2：

// 输入：
// ["CQueue","deleteHead","appendTail","appendTail","deleteHead","deleteHead"]
// [[],[],[5],[2],[],[]]
// 输出：[null,-1,null,null,5,2]
// 提示：

// 1 <= values <= 10000
// 最多会对 appendTail、deleteHead 进行 10000 次调用


type CQueue struct {
    InStack []int
    OutStack []int
}


func Constructor() CQueue {
    return CQueue{}
}


func (this *CQueue) AppendTail(value int)  {
    this.InStack = append(this.InStack, value)
}


func (this *CQueue) DeleteHead() int {
    // 需要先将 InStack 中的数据写入 OutStack
    m, n, res := len(this.InStack), len(this.OutStack), -1
    if n == 0 {
        if m == 0 {
            return res
        }
        for i := 0; i < m; i++ {
            this.OutStack = append(this.OutStack, this.InStack[m - 1 - i])
        }
        this.InStack = []int{}
    }
    res = this.OutStack[len(this.OutStack) - 1]
    this.OutStack = this.OutStack[:len(this.OutStack) - 1]
    return res
}


/**
 * Your CQueue object will be instantiated and called as such:
 * obj := Constructor();
 * obj.AppendTail(value);
 * param_2 := obj.DeleteHead();
 */