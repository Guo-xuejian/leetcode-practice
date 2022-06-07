/*
 * @lc app=leetcode.cn id=732 lang=golang
 *
 * [732] 我的日程安排表 III
 */

// @lc code=start
type MyCalendarThree struct {
	*redblacktree.Tree
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{redblacktree.NewWithIntComparator()}
}

func (t MyCalendarThree) add(x, delta int) {
	if val, ok := t.Get(x); ok {
		delta += val.(int)
	}
	t.Put(x, delta)
}

func (t MyCalendarThree) Book(start, end int) (ans int) {
	t.add(start, 1)
	t.add(end, -1)

	maxBook := 0
	for it := t.Iterator(); it.Next(); {
		maxBook += it.Value().(int)
		if maxBook > ans {
			ans = maxBook
		}
	}
	return
}

/**
 * Your MyCalendarThree object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
// @lc code=end

