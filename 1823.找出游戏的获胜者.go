/*
 * @lc app=leetcode.cn id=1823 lang=golang
 *
 * [1823] 找出游戏的获胜者
 */

// @lc code=start
func findTheWinner(n, k int) int {
	q := make([]int, n)
	for i := range q {
		q[i] = i + 1
	}
	for len(q) > 1 {
		for i := 1; i < k; i++ {
			q = append(q, q[0])[1:]
		}
		q = q[1:]
	}
	return q[0]
}

// @lc code=end

