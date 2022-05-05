/*
 * @lc app=leetcode.cn id=933 lang=golang
 *
 * [933] 最近的请求次数
 */

// @lc code=start
type RecentCounter []int

func Constructor() (_ RecentCounter) { return }

func (q *RecentCounter) Ping(t int) int {
	*q = append(*q, t)
	for (*q)[0] < t-3000 {
		*q = (*q)[1:]
	}
	return len(*q)
}

/**
 * Your RecentCounter object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Ping(t);
 */
// @lc code=end

