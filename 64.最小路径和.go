/*
 * @lc app=leetcode.cn id=64 lang=golang
 *
 * [64] 最小路径和
 */

// @lc code=start
func minPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	temp := make([][]int, m)
	for i := 0; i < m; i++ {
		temp[i] = make([]int, n)
		if i == 0 {
			copy(temp[i], grid[0])
			continue
		}
		temp[i][0] = grid[i][0]
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 && j == 0 {
				continue
			} else if i == 0 {
				temp[i][j] = grid[i][j] + temp[i][j-1]
			} else if j == 0 {
				temp[i][j] = grid[i][j] + temp[i-1][j]
			} else {
				temp[i][j] = grid[i][j] + min(temp[i-1][j], temp[i][j-1])
			}
		}
	}
	return temp[m-1][n-1]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

// @lc code=end

