/*
 * @lc app=leetcode.cn id=641 lang=golang
 *
 * [641] 设计循环双端队列
 *
 * https://leetcode.cn/problems/design-circular-deque/description/
 *
 * algorithms
 * Medium (52.99%)
 * Likes:    127
 * Dislikes: 0
 * Total Accepted:    30.4K
 * Total Submissions: 56.9K
 * Testcase Example:  '["MyCircularDeque","insertLast","insertLast","insertFront","insertFront","getRear","isFull","deleteLast","insertFront","getFront"]\n' +
  '[[3],[1],[2],[3],[4],[],[],[],[4],[]]'
 *
 * 设计实现双端队列。
 *
 * 实现 MyCircularDeque 类:
 *
 *
 * MyCircularDeque(int k) ：构造函数,双端队列最大为 k 。
 * boolean insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true ，否则返回 false 。
 * boolean insertLast() ：将一个元素添加到双端队列尾部。如果操作成功返回 true ，否则返回 false 。
 * boolean deleteFront() ：从双端队列头部删除一个元素。 如果操作成功返回 true ，否则返回 false 。
 * boolean deleteLast() ：从双端队列尾部删除一个元素。如果操作成功返回 true ，否则返回 false 。
 * int getFront() )：从双端队列头部获得一个元素。如果双端队列为空，返回 -1 。
 * int getRear() ：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1 。
 * boolean isEmpty() ：若双端队列为空，则返回 true ，否则返回 false  。
 * boolean isFull() ：若双端队列满了，则返回 true ，否则返回 false 。
 *
 *
 *
 *
 * 示例 1：
 *
 *
 * 输入
 * ["MyCircularDeque", "insertLast", "insertLast", "insertFront",
 * "insertFront", "getRear", "isFull", "deleteLast", "insertFront", "getFront"]
 * [[3], [1], [2], [3], [4], [], [], [], [4], []]
 * 输出
 * [null, true, true, true, false, 2, true, true, true, 4]
 *
 * 解释
 * MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
 * circularDeque.insertLast(1);                    // 返回 true
 * circularDeque.insertLast(2);                    // 返回 true
 * circularDeque.insertFront(3);                    // 返回 true
 * circularDeque.insertFront(4);                    // 已经满了，返回 false
 * circularDeque.getRear();                  // 返回 2
 * circularDeque.isFull();                        // 返回 true
 * circularDeque.deleteLast();                    // 返回 true
 * circularDeque.insertFront(4);                    // 返回 true
 * circularDeque.getFront();                // 返回 4
 *
 *
 *
 *
 * 提示：
 *
 *
 * 1 <= k <= 1000
 * 0 <= value <= 1000
 * insertFront, insertLast, deleteFront, deleteLast, getFront, getRear,
 * isEmpty, isFull  调用次数不大于 2000 次
 *
 *
*/

// @lc code=start
type MyCircularDeque struct {
	front, rear int
	elements    []int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{elements: make([]int, k+1)}
}

func (q *MyCircularDeque) InsertFront(value int) bool {
	if q.IsFull() {
		return false
	}
	q.front = (q.front - 1 + len(q.elements)) % len(q.elements)
	q.elements[q.front] = value
	return true
}

func (q *MyCircularDeque) InsertLast(value int) bool {
	if q.IsFull() {
		return false
	}
	q.elements[q.rear] = value
	q.rear = (q.rear + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteFront() bool {
	if q.IsEmpty() {
		return false
	}
	q.front = (q.front + 1) % len(q.elements)
	return true
}

func (q *MyCircularDeque) DeleteLast() bool {
	if q.IsEmpty() {
		return false
	}
	q.rear = (q.rear - 1 + len(q.elements)) % len(q.elements)
	return true
}

func (q MyCircularDeque) GetFront() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[q.front]
}

func (q MyCircularDeque) GetRear() int {
	if q.IsEmpty() {
		return -1
	}
	return q.elements[(q.rear-1+len(q.elements))%len(q.elements)]
}

func (q MyCircularDeque) IsEmpty() bool {
	return q.rear == q.front
}

func (q MyCircularDeque) IsFull() bool {
	return (q.rear+1)%len(q.elements) == q.front
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
// @lc code=end

