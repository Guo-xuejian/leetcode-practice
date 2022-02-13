/*
 * @lc app=leetcode.cn id=1254 lang=golang
 *
 * [1254] 统计封闭岛屿的数目
 */

// @lc code=start
func closedIsland(grid [][]int) (res int) {
	// 从边界出发
	m, n := len(grid), len(grid[0])
	directions := [][2]int{{1, 0}, {-1, 0}, {0, 1}, {0, -1}}

	var dfs func(x, y int)
	dfs = func(x, y int) { // 出得去 false，否则出不去 true(循环需要计数)
		if x < 0 || y < 0 || x >= m || y >= n || grid[x][y] == 1 {
			return
		}
		grid[x][y] = 1
		for _, direction := range directions {
			dfs(x+direction[0], y+direction[1])
		}
	}
	// 边界
	for i := 0; i < m; i++ {
		if grid[i][0] == 0 {
			dfs(i, 0)
		}
		if grid[i][n-1] == 0 {
			dfs(i, n-1)
		}
	}
	for j := 0; j < n; j++ {
		if grid[0][j] == 0 {
			dfs(0, j)
		}
		if grid[m-1][j] == 0 {
			dfs(m-1, j)
		}
	}
	for i := 1; i < m-1; i++ {
		for j := 1; j < n-1; j++ {
			if grid[i][j] == 0 { // 剩下的封闭岛的一个陆地格子
				res++
				dfs(i, j) // 处理该岛屿其余陆地格子，避免重复计数
			}
		}
	}
	return
}

// @lc code=end

