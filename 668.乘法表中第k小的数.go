/*
 * @lc app=leetcode.cn id=668 lang=golang
 *
 * [668] 乘法表中第k小的数
 */

// @lc code=start
func findKthNumber(m, n, k int) int {
	return sort.Search(m*n, func(x int) bool {
		count := x / n * n
		for i := x/n + 1; i <= m; i++ {
			count += x / i
		}
		return count >= k
	})
}

// @lc code=end

