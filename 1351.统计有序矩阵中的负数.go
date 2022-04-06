/*
 * @lc app=leetcode.cn id=1351 lang=golang
 *
 * [1351] 统计有序矩阵中的负数
 */

// @lc code=start
func countNegatives(grid [][]int) (res int) {
	m, n := len(grid), len(grid[0])
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if grid[i][j] < 0 {
				res += n - j
				break
			}
		}
	}
	return
}

// @lc code=end

