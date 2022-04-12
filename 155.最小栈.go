/*
 * @lc app=leetcode.cn id=155 lang=golang
 *
 * [155] 最小栈
 */

// @lc code=start
type MinStack struct {
	Stack    []int
	minStack []int
}

func Constructor() MinStack {
	return MinStack{}
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func (this *MinStack) Push(val int) {
	this.Stack = append(this.Stack, val)
	if len(this.minStack) > 0 {
		this.minStack = append(this.minStack, min(this.minStack[len(this.minStack)-1], val))
	} else {
		this.minStack = append(this.minStack, val)
	}
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.minStack[len(this.minStack)-1]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */
// @lc code=end

